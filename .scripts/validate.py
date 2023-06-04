import subprocess
from datetime import datetime

now = datetime.now()

# 執行 git log 命令並捕獲輸出
cmd = ['git', 'log', '--date=short', '--pretty=format:%cd']
output = subprocess.check_output(cmd).decode('utf-8').strip()

# print(output)
# 將輸出按行分割成清單
commit_dates = output.split('\n')

# 將日期字串轉換為 datetime 物件
commit_dates = [datetime.strptime(date, '%Y-%m-%d') for date in commit_dates]

# 檢查是否有連續的 commit
current_date = commit_dates[0]

if current_date.strftime('%Y-%m-%d') != now.strftime('%Y-%m-%d'):
    print("今天 %s 還沒有 Commit 喔，Let's Write Code"%now.strftime("%Y-%m-%d"))

for date in commit_dates[1:]:
    if (current_date-date).days > 1:
        print("有缺少連續的 commit 於", current_date.strftime('%Y-%m-%d'), "與", date.strftime('%Y-%m-%d'))
    current_date = date

# package main

# import (
# 	"fmt"
# 	"log"
# 	"os/exec"
# )

# func main() {
# 	out, err := exec.Command("date").Output()
# 	if err != nil {
# 		log.Fatal(err)
# 	}
# 	fmt.Printf("The date is %s\n", out)
# }
