#!/bin/sh
 
DIR='./'
IGNORED_DIRS=(".//.git" "./" ".//.scripts")
MAX_DAY=0

for DIR in $(find "$DIR" -maxdepth 1 -type d)
do
    # echo $DIR
    SKIP=false
    for IGNORED in "${IGNORED_DIRS[@]}"
    do
        if [ "$DIR" = "$IGNORED" ]
        then
            SKIP=true
            break
        fi
    done

    if [ "$SKIP" = false ]
    then
        # 使用basename命令獲取路徑的最後一個元素
        DIRNAME=$(basename "$DIR")
        # 使用cut命令剪切字串以獲取day後面的數字
        DAY=$(echo "$DIRNAME" | cut -d'-' -f2)
        # 比較當前的DAY和MAX_DAY，如果DAY大於MAX_DAY，則更新MAX_DAY
        if [ "$DAY" -gt "$MAX_DAY" ]
        then
            MAX_DAY="$DAY"
        fi
    fi
done

echo "目前最大值："$MAX_DAY
NEW_DAY=$((MAX_DAY + 1))
# echo $NEW_DAY
# 將MAX_DAY格式化為兩位數
NEW_DAY_FORMAT=$(printf "%02d" "$NEW_DAY")
# 生成資料夾名稱
NEW_DIR="Day-$NEW_DAY_FORMAT"
# 創建資料夾
mkdir "$NEW_DIR"
echo "Create Folder $NEW_DIR"
