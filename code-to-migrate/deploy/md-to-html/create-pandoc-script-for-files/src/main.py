import argparse
import logging
import os
import pathlib
import sys
import typing as tp


def get_argument_parser():
    parser = argparse.ArgumentParser(
        description="Create script to convert .md files to .html"
    )
    parser.add_argument("nginx_web_content_pathname", type=str)
    parser.add_argument("pandoc_volume_pathname", type=str)
    return parser


class Logger:
    @property
    def logger(self):
        logger = logging.getLogger(__name__)
        logger.addHandler(self._console_handler)
        logger.addHandler(self._file_handler)
        logger.setLevel(logging.DEBUG)
        return logger

    @property
    def _console_handler(self):
        c_handler = logging.StreamHandler(sys.stdout)
        c_format = logging.Formatter(self._log_format, datefmt=self._date_format)
        c_handler.setFormatter(c_format)
        return c_handler

    @property
    def _file_handler(self):
        f_handler = logging.FileHandler(filename=self._log_file_pathname, mode="w")
        f_format = logging.Formatter(self._log_format, datefmt=self._date_format)
        f_handler.setFormatter(f_format)
        return f_handler

    @property
    def _log_file_pathname(self) -> str:
        current_dir_pathname = pathlib.Path(__file__).parent.absolute()
        result = current_dir_pathname.joinpath("file.log")
        result = str(result)
        return result

    @property
    def _log_format(self) -> str:
        return "%(asctime)s - %(levelname)s - %(message)s"

    @property
    def _date_format(self) -> str:
        return "%Y-%m-%d %H:%M:%S"


logger = Logger().logger


def run(
    nginx_web_content_pathname: str,
    pandoc_volume_pathname: str,
):
    script_to_create_pathname = str(pathlib.PurePath(pandoc_volume_pathname, "run-on-files-convert-md-to-html"))
    pandoc_script_convert_md_to_html_file_pathname = str(pathlib.PurePath(pandoc_volume_pathname, "convert-md-to-html"))
    logger.debug(f"Init export file {script_to_create_pathname}")
    # TODO move constants to config.py
    md_pathnames_to_convert_file_pathname = "/tmp/path-names-to-convert.txt"
    md_pathnames_converted_file_pathname = "/tmp/path-names-converted.txt"
    root_relative_pathnames_file_pathname = "/tmp/root-relative-pathnames.txt"
    output_directory_pathname = nginx_web_content_pathname
    export_to_file_the_md_pathnames_to_convert(
        nginx_web_content_pathname, md_pathnames_to_convert_file_pathname
    )
    export_to_file_the_html_pathnames_converted(
        nginx_web_content_pathname,
        md_pathnames_to_convert_file_pathname,
        output_directory_pathname=output_directory_pathname,
        result_file_pathname=md_pathnames_converted_file_pathname,
    )
    export_to_file_the_root_relative_pathnames(
        nginx_web_content_pathname,
        md_pathnames_to_convert_file_pathname,
        result_file_pathname=root_relative_pathnames_file_pathname,
    )
    export_to_file_the_script_combine_files(
        pandoc_script_convert_md_to_html_file_pathname,
        md_pathnames_to_convert_file_pathname,
        md_pathnames_converted_file_pathname,
        root_relative_pathnames_file_pathname,
        script_to_create_pathname,
        pandoc_volume_pathname,
    )


def export_to_file_the_html_pathnames_converted(
    analized_directory_pathname: str,
    md_pathnames_to_convert_file_pathname: str,
    output_directory_pathname: str,
    result_file_pathname: str,
):
    logger.debug(f"Init export file {result_file_pathname}")
    with open(md_pathnames_to_convert_file_pathname, "r") as f_to_read, open(
        result_file_pathname, "w"
    ) as f_to_write:
        for file_to_convert_pathname in f_to_read.read().splitlines():
            logger.debug(f"Pathname to convert: {file_to_convert_pathname}")
            pathname_converted = get_pathname_converted(
                analized_directory_pathname,
                file_to_convert_pathname,
                output_directory_pathname,
            )
            logger.debug(f"Pathname converted: {pathname_converted}")
            f_to_write.write(pathname_converted)
            f_to_write.write("\n")


def export_to_file_the_root_relative_pathnames(
    nginx_web_content_pathname: str,
    md_pathnames_to_convert_file_pathname: str,
    result_file_pathname: str,
):
    logger.debug(f"Init export file {result_file_pathname}")
    nginx_web_content_path = pathlib.PurePath(nginx_web_content_pathname)
    root_path_detector = RootPathDetector(nginx_web_content_path)
    with open(md_pathnames_to_convert_file_pathname, "r") as f_to_read, open(
        result_file_pathname, "w"
    ) as f_to_write:
        for file_to_convert_pathname in f_to_read.read().splitlines():
            logger.debug(f"Pathname to convert: {file_to_convert_pathname}")
            file_to_convert_path = pathlib.PurePath(file_to_convert_pathname)
            root_relative_pathname = (
                root_path_detector.get_root_relative_pathname_from_file_path(
                    file_to_convert_path
                )
            )
            logger.debug(f"Root relative pathname: {root_relative_pathname}")
            f_to_write.write(root_relative_pathname)
            f_to_write.write("\n")


def export_to_file_the_script_combine_files(
    pandoc_script_convert_md_to_html_file_pathname: str,
    md_pathnames_to_convert_file_pathname: str,
    md_pathnames_converted_file_pathname: str,
    root_relative_pathnames_file_pathname: str,
    result_file_pathname: str,
    pandoc_volume_pathname: str,
):
    with open(md_pathnames_to_convert_file_pathname) as to_convert_file, open(
        md_pathnames_converted_file_pathname
    ) as converted_file, open(root_relative_pathnames_file_pathname) as root_file, open(
        result_file_pathname, "w"
    ) as script_file:
        to_convert_lines = to_convert_file.read().splitlines()
        converted_lines = converted_file.read().splitlines()
        root_lines = root_file.read().splitlines()
        assert len(to_convert_lines) == len(converted_lines) == len(root_lines)
        for file_to_convert_pathname, file_converted_pathname, root_directory_relative_path_name in zip(
            to_convert_lines, converted_lines, root_lines
        ):
            command = "/bin/sh {} {} {} {} {}".format(
                pandoc_script_convert_md_to_html_file_pathname,
                file_to_convert_pathname,
                file_converted_pathname,
                root_directory_relative_path_name,
                pandoc_volume_pathname,
            )
            logger.debug(f"Command: {command}")
            script_file.write(command)
            script_file.write("\n")


class DirectoryAnalyzer:
    def get_md_pathnames(self, pathname: str) -> tp.Iterator[str]:
        logger.debug(f"Init checking {pathname}")
        for (dir_pathname, dirnames, filenames) in os.walk(
            pathname, onerror=self._os_walk_exception_handler
        ):
            # print(dir_pathname, dirnames, filenames)
            for filename in filenames:
                if self._is_md_file(filename):
                    yield str(pathlib.PurePath(dir_pathname, filename))

    def _os_walk_exception_handler(self, exception_instance):
        raise exception_instance

    @staticmethod
    def _is_md_file(filename: str) -> bool:
        return pathlib.PurePath(filename).suffix.lower().strip() == ".md"


class RootPathDetector:
    def __init__(self, nginx_web_content_path: pathlib.PurePath):
        self._nginx_web_content_path = nginx_web_content_path

    def get_root_relative_pathname_from_file_path(
        self, file_path: pathlib.PurePath
    ) -> str:
        file_path_without_filename = file_path.parent
        return (
            "."
            if self._nginx_web_content_path == file_path_without_filename
            else self._get_root_relative_pathname_when_file_with_different_path(
                file_path_without_filename,
            )
        )

    def _get_root_relative_pathname_when_file_with_different_path(
        self,
        file_path_without_filename: pathlib.PurePath,
    ) -> str:
        folders_between_files_path: pathlib.PurePath = (
            file_path_without_filename.relative_to(self._nginx_web_content_path)
        )
        folders_between_files = str(folders_between_files_path).split("/")
        relative_pathnames = [".." for _ in folders_between_files]
        result = "/".join(relative_pathnames)
        return result


def get_path_substract_common_parts(
    path_1: pathlib.PurePath, path_2: pathlib.PurePath
) -> pathlib.PurePath:
    return path_1.relative_to(path_2)


def export_to_file_the_md_pathnames_to_convert(
    nginx_web_content_pathname: str,
    result_file_pathname: str,
):
    logger.debug(f"Init export file {result_file_pathname}")
    with open(result_file_pathname, "w") as f:
        for md_pathname in DirectoryAnalyzer().get_md_pathnames(nginx_web_content_pathname):
            logger.debug(f"Detected .md file: {md_pathname}")
            f.write(md_pathname)
            f.write("\n")


def get_pathname_converted(
    analized_directory_pathname: str,
    file_to_convert_pathname: str,
    output_directory_pathname: str,
) -> str:
    path_to_convert = pathlib.PurePath(file_to_convert_pathname)
    path_to_convert_without_analized_path = get_path_substract_common_parts(
        path_to_convert, pathlib.PurePath(analized_directory_pathname)
    )
    filename_to_convert = path_to_convert.name
    filename_converted = get_filename_set_extension(filename_to_convert, ".html")
    path_converted_without_analized_path = (
        path_to_convert_without_analized_path.with_name(filename_converted)
    )
    path_converted = pathlib.PurePath(output_directory_pathname).joinpath(
        path_converted_without_analized_path,
    )
    pathname_converted = str(path_converted)
    return pathname_converted


def get_filename_set_extension(filename: str, extension: str) -> str:
    return pathlib.PurePath(filename).with_suffix(extension).name


if __name__ == "__main__":
    args = get_argument_parser().parse_args()
    run(
        args.nginx_web_content_pathname,
        args.pandoc_volume_pathname,
    )
