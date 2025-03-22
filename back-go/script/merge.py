import subprocess
import os


commend = ['ffmpeg -i \"', '', '\" -i \"', '', '\" -c:v copy -c:a copy -map 0:v:0 -map 1:a:0 "', '', '\"']


def merge( videoFileName, audioFileName, finalFileName ):
        command = [
            'ffmpeg',
            '-i', videoFileName,
            '-i', audioFileName,
            '-c:v', 'copy',
            '-c:a', 'copy',
            '-map', '0:v:0',
            '-map', '1:a:0',
            finalFileName
        ]

        try:
            subprocess.run(command, check=True)
            print("Success")
        except subprocess.CalledProcessError as e:
            print(f"Error Merge" + str(e))


if __name__ == "__main__":
    merge(
         "E:\\project\\Yuyu-std.github.io\\back\\script\\tem\\抽象复合函数求偏导，恼人的下标怎样理解？其实很简单_哔哩哔哩_bilibili.mp4",
         "E:\\project\\Yuyu-std.github.io\\back\\script\\tem\\抽象复合函数求偏导，恼人的下标怎样理解？其实很简单_哔哩哔哩_bilibili (1).mp4",
         "E:\\project\\Yuyu-std.github.io\\back\\script\\tem\\test.mp4")