- Learn Vim Progressively <http://yannesposito.com/Scratch/en/blog/Learn-Vim-Progressively/>
- Learn Vim Progressively 翻译 <http://www.oschina.net/translate/learn-vim-progressively>
- 图片教程 <http://www.viemu.com/a_vi_vim_graphical_cheat_sheet_tutorial.html> `所有文件\dev\vi`

# 常用

- 保留每一行的前10个字符 `qq010lDjq` + `9999@q` 或 `:%s/^\(.\{10}\).*$/\1/g`
- 删除每一行前10个字符 `:%s/^.\{10\}//` 或 `:%s/^\(.\{9}\).*$/\1/g`
