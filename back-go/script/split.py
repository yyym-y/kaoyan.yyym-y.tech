import os
import re
import uuid
import subprocess

def get_all_mp4_files(root_folder):
    mp4_files = []
    mp4_name = []
    firstDirName = []
    ppos = -1
    for dirpath, dirnames, filenames in os.walk(root_folder):
        if len(firstDirName) == 0:
            firstDirName = dirnames
        for pos in range(len(filenames)):
            filename = filenames[pos]
            if filename.endswith('.mp4'):
                full_path = os.path.join(dirpath, filename)
                mp4_files.append(full_path)
                if(ppos != -1):
                    name = filename.replace(".mp4", "").replace("语法快速入门", "").replace("阅读命题句每日一句", "").replace("8小时带写64个句子", "")
                    mp4_name.append(firstDirName[ppos] + " - " + name)
        ppos += 1
    # combined = list(zip(mp4_files, mp4_name))
    # combined_sorted = sorted(combined, key=lambda x: int(re.match(r'^\d+', x[1]).group()))
    # sorted_mp4_files, sorted_mp4_name = zip(*combined_sorted)
    # sorted_mp4_files = list(sorted_mp4_files)
    # sorted_mp4_name = list(sorted_mp4_name)


    
    return mp4_files, mp4_name

def convert_mp4_to_m3u8(input_folder, begin_fold_num, vedio_id):
    # 获取所有的mp4文件
    mp4_files, vedio_name = get_all_mp4_files(input_folder)
    
    # 创建输出文件夹
    output_folder = os.path.join(input_folder, "output")
    os.makedirs(output_folder, exist_ok=True)

    for index, mp4_file in enumerate(mp4_files, start=1):
        # 创建以数字命名的子文件夹
        subfolder_name = str(index + begin_fold_num - 1)
        subfolder_path = os.path.join(output_folder, subfolder_name)
        os.makedirs(subfolder_path, exist_ok=True)

        # 输入和输出文件路径
        input_file_path = mp4_file

        m3u8_file_name = str(uuid.uuid4()).replace("-", "")
        m3u8_file_path = os.path.join(subfolder_path, f"{m3u8_file_name}.m3u8")

        # FFmpeg命令
        command = [
            'ffmpeg',
            '-i', input_file_path,
            '-c', 'copy',
            # '-c:v', 'libx264',  # 使用H.264编码器
            # '-c:a', 'aac',      # 使用AAC音频编码
            '-start_number', '0',
            '-hls_time', '10',
            '-hls_list_size', '0',
            '-f', 'hls',
            m3u8_file_path
        ]

        # 执行FFmpeg命令
        try:
            subprocess.run(command, check=True)
            with open("./insert_m3u8.sql", 'a', encoding='utf-8') as file:
                sql_commend = [
                    'INSERT INTO playinfo VALUES (', str(vedio_id) , ', ',
                    str(index + begin_fold_num - 1), ', 0, \"',
                    m3u8_file_name + ".m3u8\", \"",
                    vedio_name[index - 1], "\");\n\n"
                ]
                file.write("".join(sql_commend))
        except subprocess.CalledProcessError as e:
            print(f"Error converting {mp4_file}: {e}")

if __name__ == "__main__":
    input_folder = "E:\\BaiduNetdiskDownload\\all\\04.基础阶段\\01.基础起步\\英语基础起步"
    # list1, list2 = get_all_mp4_files(input_folder)
    # with open("./kk.txt", "w", encoding='utf-8') as ff:
    #     for pr in list2:
    #         ff.write(pr + "\n")
    #     for pr in list1:
    #         ff.write(pr + "\n")
    convert_mp4_to_m3u8(input_folder, 1, 26)
    # print(str(uuid.uuid4()).replace("-", "") + ".m3u8")