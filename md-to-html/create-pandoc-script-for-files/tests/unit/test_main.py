import pathlib
import unittest

from src import main


class TestDirectoryAnalyzer(unittest.TestCase):
    def test_is_md_file_is_true(self):
        self.assertTrue(main.DirectoryAnalyzer._is_md_file("foo.md"))

    def test_is_md_file_is_true_if_upper_case_extension(self):
        self.assertTrue(main.DirectoryAnalyzer._is_md_file("foo.MD"))

    def test_is_md_file_is_true_if_trailing_space(self):
        self.assertTrue(main.DirectoryAnalyzer._is_md_file("foo.md "))

    def test_is_md_file_is_false(self):
        self.assertFalse(main.DirectoryAnalyzer._is_md_file("foo.html"))


class TestFunctions(unittest.TestCase):
    def test_get_convered_pathname(self):
        result = main.get_pathname_converted(
            analized_directory_pathname="/home/files",
            file_to_convert_pathname="/home/files/folder/foo.html",
            output_directory_pathname="/tmp/html",
        )
        self.assertEqual("/tmp/html/folder/foo.html", result)

    def test_get_filename_set_extension(self):
        filename = "foo.md"
        result = main.get_filename_set_extension(filename, ".html")
        self.assertEqual("foo.html", result)
        self.assertTrue(isinstance(result, str))


class TestRootPathDetector(unittest.TestCase):
    def test_get_root_relative_pathname_from_file_path(self):
        nginx_web_content_path = pathlib.PurePath("/foo/bar")
        html_path = pathlib.PurePath("/foo/bar/folder_1/folder_2/index.html")
        root_path_detector = main.RootPathDetector(nginx_web_content_path)
        self.assertEqual(
            "../..",
            root_path_detector.get_root_relative_pathname_from_file_path(html_path),
        )

    def test_get_root_relative_pathname_from_file_path_if_files_in_same_path(self):
        nginx_web_content_path = pathlib.PurePath("/foo/bar")
        html_path = pathlib.PurePath("/foo/bar/index.html")
        root_path_detector = main.RootPathDetector(nginx_web_content_path)
        self.assertEqual(
            ".",
            root_path_detector.get_root_relative_pathname_from_file_path(html_path),
        )
