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
    void BgBlack() *Color{
        if !this.disabled {
            cout << BG_BLACK;
        }
    };
    void BgBlue() *Color{
        if !this.disabled {
            cout << BG_BLACK;
        }
    };
    void BgCyan() *Color{
        if !this.disabled {
            cout << BG_BLACK;
        }
    };
    void BgGreen() *Color{
         if !this.disabled {
             cout << BG_BLACK;
         }
     };
    void BgMagenta() *Color{
        if !this.disabled {
           cout << BG_BLACK;
        }
    };
    void BgRed() *Color{
       if !this.disabled {
           cout << BG_BLACK;
       }
    };
    void BgWhite() *Color{
        if !this.disabled {
            cout << BG_BLACK;
        }
    };
    void BgYellow() *Color{
        if !this.disabled {
            cout << BG_BLACK;
        }
    };

    void Black() *Color{
        if !this.disabled {
            cout << BG_BLACK;
        }
    };
    void Blue() *Color{
        if !this.disabled {
            cout << BG_BLACK;
        }
    };
    void Cyan() *Color{
        if !this.disabled {
        cout << BG_BLACK;
        }
    };
    void Green() *Color{
        if !this.disabled {
        cout << BG_BLACK;
        }
    };
    void Magenta() *Color{
        if !this.disabled {
            cout << BG_BLACK;
        }
    };
    void Red() *Color{
        if !this.disabled {
            cout << BG_BLACK;
        }
    };
    void White() *Color{
        if !this.disabled {
            cout << BG_BLACK;
        }
    };
    void Yellow() *Color{
        if !this.disabled {
            cout << BG_BLACK;
        }
    };

    void Reset() *Color{
        if !this.disabled {
            cout << RESET;
        }
    };
    void Clear() *Color{
        if !this.disabled {
            cout << CLEAR;
        }
    };
    void Bold() *Color{
        if !this.disabled {
            cout << BOLD;
        }
    };
    void Dim() *Color{
        if !this.disabled{
            cout << DIM;
        }
    };
    void Underline() *Color{
        if !this.disabled{
            cout << UNDERLINE;
        }
    };
    void Blink() *Color{
        if !this.disabled{
            cout << BLINK;
        }
    };
    void Reverse() *Color{
        if !this.disabled{
            cout << REVERSE;
        }
    };
    void HiddenText() *Color{
        if !this.disabled{
            cout << HIDDEN_TEXT;
        }
    };
    void Strikethrough() *Color{
        if !this.disabled{
            cout << STRIKETHROUGH;
        }
    };

    void Print(string msg) *Color{
        cout << msg;
    };
    void Print(string *msg) *Color{
        cout << *msg;
    };
    void Printf(string fmt, string msg) *Color{
        cout << *msg;
    };
    void Printf(string *fmt, string *msg) *Color{
        cout << *msg;
    };
    void Println(string msg) *Color{
        cout << *msg;
    };
    void Println(string *msg) *Color{
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
