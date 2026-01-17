from distutils import dir_util
import pathlib
import unittest

from src import main

# TODO remove generated files when the tests ends:
# ```
# ./drop-files-created-by-the-tests
# ```


class TestDirectoryAnalyzer(unittest.TestCase):
    def setUp(self):
        script_dir = pathlib.Path(__file__).parent.absolute()
        tests_dir = script_dir.parent
        self.nginx_web_content_pathname = str(
            pathlib.PurePath(tests_dir, "files", "fake-project")
        )

    def test_get_md_pathnames(self):
        md_files = [
            pathname
            for pathname in main.DirectoryAnalyzer().get_md_pathnames(
                self.nginx_web_content_pathname
            )
        ]
        self.assertEqual(
            [
                str(pathlib.PurePath(self.nginx_web_content_pathname, "foo.md")),
                str(pathlib.PurePath(self.nginx_web_content_pathname, "folder-1", "bar.md")),
            ],
            md_files,
        )

    def test_exception_is_raised_if_nginx_web_content_pathname_does_not_exist(self):
        nginx_web_content_pathname = "/foo/bar/asdf/foo/bar"
        with self.assertRaises(FileNotFoundError):
            for _ in main.DirectoryAnalyzer().get_md_pathnames(nginx_web_content_pathname):
                pass


class TestFunctions(unittest.TestCase):
    def setUp(self):
        script_dir = pathlib.Path(__file__).parent.absolute()
        tests_dir = script_dir.parent
        self.test_files_path = pathlib.PurePath(tests_dir, "files")
        self.test_fake_project_pathname = str(
            self.test_files_path.joinpath("fake-project")
        )
        self.test_md_pathnames_to_convert_file_pathname = str(
            self.test_files_path.joinpath("pathnames-to-convert.txt")
        )

    def test_export_to_file_the_md_pathnames_to_convert(self):
        result_file_pathname = "/tmp/input.txt"
        main.export_to_file_the_md_pathnames_to_convert(
            nginx_web_content_pathname=self.test_fake_project_pathname,
            result_file_pathname=result_file_pathname,
        )
        with open(result_file_pathname, "r") as f:
            result = f.read()
        expected_result = "{}\n{}\n".format(
            f"{self.test_fake_project_pathname}/foo.md",
            f"{self.test_fake_project_pathname}/folder-1/bar.md",
        )
        self.assertEqual(expected_result, result)

    def test_export_to_file_the_html_pathnames_converted(self):
        result_file_pathname = "/tmp/output.txt"
        output_directory_pathname = "/tmp/html"
        main.export_to_file_the_html_pathnames_converted(
            analized_directory_pathname="/home/files",
            md_pathnames_to_convert_file_pathname=self.test_md_pathnames_to_convert_file_pathname,
            output_directory_pathname=output_directory_pathname,
            result_file_pathname=result_file_pathname,
        )
        with open(result_file_pathname, "r") as f:
            result = f.read()
        expected_result = "{}\n{}\n".format(
            f"{output_directory_pathname}/foo.html",
            f"{output_directory_pathname}/folder-1/bar.html",
        )
        self.assertEqual(expected_result, result)

    def test_export_to_file_the_root_relative_pathnames(self):
        nginx_web_content_pathname = "/home"
        result_file_pathname = "/tmp/output-relative.txt"
        main.export_to_file_the_root_relative_pathnames(
            nginx_web_content_pathname=nginx_web_content_pathname,
            md_pathnames_to_convert_file_pathname=self.test_md_pathnames_to_convert_file_pathname,
            result_file_pathname=result_file_pathname,
        )
        with open(result_file_pathname, "r") as f:
            result = f.read()
        expected_result = "{}\n{}\n".format(
            "..",
            "../..",
        )
        self.assertEqual(expected_result, result)


class TestRun(unittest.TestCase):
    def setUp(self):
        self.maxDiff = None  # Show complete diff when test fails.
        script_dir = pathlib.Path(__file__).parent.absolute()
        tests_dir = script_dir.parent
        self.nginx_web_content_pathname = "/tmp/cmoli.es/html"
        pathname_with_files_to_analyze = str(
            pathlib.PurePath(tests_dir, "files", "fake-project")
        )
        dir_util.copy_tree(pathname_with_files_to_analyze, self.nginx_web_content_pathname)

    def test_run(self):
        pandoc_volume_pathname = "/tmp"
        script_to_create_pathname=f"{pandoc_volume_pathname}/run-on-files-convert-md-to-html"
        main.run(
            nginx_web_content_pathname=self.nginx_web_content_pathname,
            pandoc_volume_pathname=pandoc_volume_pathname,
        )
        with open(script_to_create_pathname, "r") as f:
            result = f.read()
        expected_result = "{}\n{}\n".format(
            "/bin/sh /tmp/convert-md-to-html /tmp/cmoli.es/html/foo.md /tmp/cmoli.es/html/foo.html . /tmp",
            "/bin/sh /tmp/convert-md-to-html /tmp/cmoli.es/html/folder-1/bar.md /tmp/cmoli.es/html/folder-1/bar.html .. /tmp",
        )
        self.assertEqual(expected_result, result)
