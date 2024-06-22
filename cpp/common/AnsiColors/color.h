/**
 * @name AnsiColors/color.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include <stdio>

using namespace std;

class Color{
private:
    bool disabled;
public:
    Color(bool disabled=false){
        this.disabled=disabled;
    }
    ~Color(){}

    void inline UnixTime() *Color {
        if !this.disabled {
            auto now = std::chrono::system_clock::now();
            std::time_t timestamp = std::chrono::system_clock::to_time_t(now);
            std::cout << timestamp << std::endl;
        }
    };
    void inline UnixTimeNano() {
        if (!this->disabled) {
            auto now = std::chrono::system_clock::now();
            auto duration = now.time_since_epoch();
            auto nanoseconds = std::chrono::duration_cast<std::chrono::nanoseconds>(duration).count();
            std::cout << nanoseconds << std::endl;
        }
    }
    void inline BgBlack() *Color{
        if !this.disabled {
            cout << BG_BLACK;
        }
    };
    void inline BgBlue() *Color{
        if !this.disabled {
            cout << BG_BLACK;
        }
    };
    void inline BgCyan() *Color{
        if !this.disabled {
            cout << BG_BLACK;
        }
    };
    void inline BgGreen() *Color{
         if !this.disabled {
             cout << BG_BLACK;
         }
     };
    void inline BgMagenta() *Color{
        if !this.disabled {
           cout << BG_BLACK;
        }
    };
    void inline BgRed() *Color{
       if !this.disabled {
           cout << BG_BLACK;
       }
    };
    void inline BgWhite() *Color{
        if !this.disabled {
            cout << BG_BLACK;
        }
    };
    void inline BgYellow() *Color{
        if !this.disabled {
            cout << BG_BLACK;
        }
    };

    void inline Black() *Color{
        if !this.disabled {
            cout << BG_BLACK;
        }
    };
    void inline Blue() *Color{
        if !this.disabled {
            cout << BG_BLACK;
        }
    };
    void inline Cyan() *Color{
        if !this.disabled {
        cout << BG_BLACK;
        }
    };
    void inline Green() *Color{
        if !this.disabled {
        cout << BG_BLACK;
        }
    };
    void inline Magenta() *Color{
        if !this.disabled {
            cout << BG_BLACK;
        }
    };
    void inline Red() *Color{
        if !this.disabled {
            cout << BG_BLACK;
        }
    };
    void inline White() *Color{
        if !this.disabled {
            cout << BG_BLACK;
        }
    };
    void inline Yellow() *Color{
        if !this.disabled {
            cout << BG_BLACK;
        }
    };

    void inline Reset() *Color{
        if !this.disabled {
            cout << RESET;
        }
    };
    void inline Clear() *Color{
        if !this.disabled {
            cout << CLEAR;
        }
    };
    void inline Bold() *Color{
        if !this.disabled {
            cout << BOLD;
        }
    };
    void inline Dim() *Color{
        if !this.disabled{
            cout << DIM;
        }
    };
    void inline Underline() *Color{
        if !this.disabled{
            cout << UNDERLINE;
        }
    };
    void inline Blink() *Color{
        if !this.disabled{
            cout << BLINK;
        }
    };
    void inline Reverse() *Color{
        if !this.disabled{
            cout << REVERSE;
        }
    };
    void inline HiddenText() *Color{
        if !this.disabled{
            cout << HIDDEN_TEXT;
        }
    };
    void inline Strikethrough() *Color{
        if !this.disabled{
            cout << STRIKETHROUGH;
        }
    };

    void inline Print(string msg) *Color{
        cout << msg;
    };
    void inline Print(string *msg) *Color{
        cout << *msg;
    };
    void inline Printf(string fmt, string msg) *Color{
        cout << *msg;
    };
    void inline Printf(string *fmt, string *msg) *Color{
        cout << *msg;
    };
    void inline Println(string msg) *Color{
        cout << *msg;
    };
    void inline Println(string *msg) *Color{
        cout << *msg;
    };
}
/*
 * Background Colors
 */
void inline BgBlack() *Color {
    Color c
	c.BgBlack()
	return &c
};
void inline BgBlue() *Color {
    Color c
    c.BgBlue()
    return &c
};
void inline BgCyan() *Color {
    Color c
    c.BgCyan()
    return &c
};
void inline BgGreen() *Color {
    Color c
    c.BgGreen()
    return &c
};
void inline BgMagenta() *Color {
    Color c
    c.BgMagenta()
    return &c
};
void inline BgRed() *Color {
    Color c
    c.BgRed()
    return &c
};
void inline BgWhite() *Color {
    Color c
    c.BgWhite()
    return &c
};
void inline BgYellow() *Color {
    Color c
    c.BgYellow()
    return &c
};
/*
 * Foreground Colors
 */
void inline Black() *Color {
    Color c
    c.Black()
    return &c
};
void inline Blue() *Color {
    Color c
    c.Blue()
    return &c
};
void inline Cyan() *Color {
    Color c
    c.Cyan()
    return &c
};
void inline Green() *Color {
    Color c
    c.Green()
    return &c
};
void inline Magenta() *Color {
    Color c
    c.Magenta()
    return &c
};
void inline Red() *Color {
    Color c
    c.Red()
    return &c
};
void inline White() *Color {
    Color c
    c.White()
    return &c
};
void inline Yellow() *Color {
    Color c
    c.Yellow()
    return &c
};
/*
 * Controls
 */
void inline Reset() *Color {
    Color c
    c.Reset()
    return &c
};
void inline Clear() *Color  {
    Color c
    c.Clear()
    return &c
};
void inline Bold() *Color{
    Color c
    c.Bold()
    return &c
};
void inline Dim() *Color{
    Color c
    c.Dim()
    return &c
};
void inline Underline() *Color{
    Color c
    c.Underline()
    return &c
};
void inline Blink() *Color{
    Color c
    c.Blink()
    return &c
};
void inline Reverse() *Color{
    Color c
    c.Reverse()
    return &c
};
void inline HiddenText() *Color{
    Color c
    c.HiddenText()
    return &c
};
void inline Strikethrough() *Color{
    Color c
    c.Strikethrough()
    return &c
};
/*
 * Print Commands
 */
void inline Print(string msg) *Color{
    Color c
    c.Print(&msg)
    return &c
};
void inline Print(string *msg) *Color{
    Color c
    c.Print(msg)
    return &c
};
void inline Printf(string msg) *Color{
    Color c
    c.Printf(&msg)
    return &c
};
void inline Printf(string *msg) *Color{
    Color c
    c.Printf(msg)
    return &c
};
void inline Println(string msg) *Color{
    Color c
    c.Println(&msg)
    return &c
};
void inline Println(string *msg) *Color{
    Color c
    c.Println(msg)
    return &c
};
