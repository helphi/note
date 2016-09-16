# 多窗口

## 打开窗口

ex command | vi command | Description
---|---|---
:[n]split [++opt] [+cmd] [file] | ^Ws ^WS ^W^S | Split the current window into two from side to side, placing <br> the cursor in the new window. The optional file argument <br> places that file in the newly created window. The windows <br> are created as equal in size as possible, determined by free <br> window space.
:[n]new [++opt] [+cmd] | ^Wn ^W^N | Same as :split, but start the new window editing an empty <br> file. Note that the buffer will have no name until one is <br> assigned.
:[n]sview [++opt] [+cmd] [file] | | Read-only version of :split.
:[n]sfind [++opt] [+cmd] [file] | | Split window and open file (if specified) in the new window. <br> Look for file in the path.
:[n]vsplit [++opt] [+cmd] [file] | ^Wv ^W^V | Split current window into two from top to bottom and open <br> file (if specified) in the new window.
:[n]vnew [++opt] [+cmd] | | Vertical version of :new.

```bash
:split #Split window.
:vsplit #Split window vertically.
:split otherfile #Split window and open aother file.
:[n]split [++opt] [+cmd] [file] 
:15split ++fileformat=unix otherfile #Set the new window to be 15 lines tall; open file with unix file format.
:[n]new [++opt] [+cmd] [file] #In addition to creating the new window, the WinLeave, WinEnter, BufLeave, and BufEnter autocommands execute.
:vnew
:sview filename #Splits the screen horizontally to open a new window and sets the readonly for that buffer. :sview requires the filename argument.
:sfind [++opt] [+cmd] filename #Works like :split, but looks for the filename in the path. If Vim does not find the file, it doesn’t split the window.
:topleft [cmd] #tells Vim to execute cmd(split, new ...) and display a new window with the cursor at the top left if cmd opens a new file
```

# 在窗口间移动

Command | Description
---|---
^W<DOWN> <br> ^W^J <br> ^Wj | Move to the next window down. <br>Note that this command does not cycle through the windows; it simply <br>moves down to the next window below the current window. If the cursor is <br>in a window at the bottom of the screen, this command has no effect. Also, <br>this command bypasses adjacent windows on its “way down”; for example, <br>if there is a window to the right of the current window, the command does <br>not jump across to the adjacent window. (Use ^W^W to cycle <br> through windows.) 
^W<UP> <br> ^W^K <br> ^Wk | Move to the next window up. This is the opposite-direction counterpart to the ^W j command. 
^W<LEFT> <br> ^W^H | Move to the window to the left of the current window.
^Wh <br> ^W<BS> <br> ^W<RIGHT> <br> ^W^L <br> ^Wl | Move to the window to the right of the current window. 
^Ww <br> ^W^W | Move to the next window below or to the right. Note that this command, unlike ^Wj , will cycle through all of the Vim windows. When the <br>lowermost window is reached, Vim restarts the cycle and moves to the top leftmost window.
^W | Move to next window above or to the left. This is the upward counterpart to the ^W w command.
^Wt <br> ^W^T | Move cursor to the top leftmost window.
^Wb <br> ^W^B | Move cursor to the bottom rightmost window.
^Wp <br> ^W^P | Move to the previous (last accessed) window.

# 移动窗口
Command | Description
---|---
^Wr <br> ^W^R | Rotate windows down or to the right. 
^WR | Rotate windows up or to the left.
^Wx <br> ^W^X | Swap positions with the next window, or if issued with a count n, swap with nth next window.
^WK | Move window to top of screen and use full width. The cursor stays with the moved window.
^WJ | Move window to bottom of screen and use full width. The cursor stays with the moved window.
^WH | Move window to left of screen and use full height. The cursor stays with the moved window.
^WL | Move window to right of screen and use full height. The cursor stays with the moved window.
^WT | Move window to new tab. The cursor stays with the moved window. If the current window is the only window in the current tab, no action is taken.

# 改变窗口大小

Command | Description
---|---
^W= | Resize all windows equally. The current window honors the settings of the winheight and winwidth options.
:resize -n <br> ^W- | Decrease the current window size. The default amount is one line.
:resize +n <br> ^W+ | Increase the current window size. The default amount is one line.
:resize n <br> ^W^_ <br> ^W_ | Set the current window height. The default is to maximize window height (unless n is specified).
zn <ENTER> | Set the current window height to n.
^W< | Increase the current window width. The default amount is one column.
^W> | Decrease the current window width. The default amount is one column.
:vertical resize n <br> ^W | Set the current window width to n. The default is to make window as wide as possible.