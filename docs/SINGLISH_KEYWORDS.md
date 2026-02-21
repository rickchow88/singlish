# SINGLISH Keyword Reference

This document provides a comprehensive reference for all **SINGLISH** keywords — their Go equivalents, cultural origin, and usage context. Singlish is built on top of Go; every Singlish keyword maps 1:1 to a Go keyword or built-in.

## 1. Design Philosophy

- **Authenticity**: Keywords should be terms actually heard in Singaporean kopitiams, offices, and streets (e.g., `chiong`, `dapao`, `shiok`).
- **Intuition**: The Singlish term should map somewhat logically to the programming concept (e.g., `dapao` for `import` implies "bringing something back to use").
- **Humor**: Maintaining the "cursed" spirit with tongue-in-cheek Singlish adaptations.

## 2. Full Keyword Mapping

### 2.1 Top-Level Structure

| Concept | Standard (Go) | CURSED | **Proposed SINGLISH** | Reason / Context |
|:--- |:--- |:--- |:--- |:--- |
| **Package** | `package` | `vibe` | **`kampung`** | A package is a community/village. Code lives in a *kampung*. |
| **Import** | `import` | `yeet` | **`dapao`** | To take away / pack. You existing libraries to *dapao* them into your code. |
| **Function** | `func` | `slay` | **`action`** | "Eh, see him *action* only." Defining an action to perform. |
| **Main** | `main` | `main` | **`boss`** | The main boss function. Entry point. |

### 2.2 Variables & Constants

| Concept | Standard (Go) | CURSED | **Proposed SINGLISH** | Reason / Context |
|:--- |:--- |:--- |:--- |:--- |
| **Variable** | `var` | `sus` | **`got`** | "Got [name] [type]". implies existence. |
| **Constant** | `const` | `facts` | **`confirm`** | "Confirm plus chop", it won't change. |
| **Iota** | `iota` | `count` | **`auto`** | Automatic counter for constants. |
| **Type Def** | `type` | `be_like` | **`pattern`** | "More pattern than badminton". Defining the pattern of data. |

### 2.3 Control Flow

| Concept | Standard (Go) | CURSED | **Proposed SINGLISH** | Reason / Context |
|:--- |:--- |:--- |:--- |:--- |
| **If** | `if` | `ready` | **`nasi`** | Teochew for "If". "Nasi you want...". |
| **Else** | `else` | `otherwise` | **`den`** | "Kalo this... *den* that". (Logic flow: If -> Then/Else). |
| **Select** | `select` | `vibez` | **`tikam`** | Random selection / gamble. Perfect for `select`. |
| **Switch** | `switch` | `vibe_check` | **`see_how`** | "See how things are", then decide. |
| **Case** | `case` | `mood` | **`say`** | "Let's *say* it's this...". |
| **Fallthrough** | `fallthrough` | `yeet_down` | **`tompang`** | "Hitch a ride" to the next case. |
| **Default** | `default` | `basic` | **`anyhow`** | "Anyhow do". The fallback option. |
| **Goto** | `goto` | `zoom` | **`flykite`** | "Flykite to". Skip/Jump to label. |

### 2.4 Loops

| Concept | Standard (Go) | CURSED | **Proposed SINGLISH** | Reason / Context |
|:--- |:--- |:--- |:--- |:--- |
| **For** | `for` | `bestie` | **`loop`** | Simple. Can also use **`again`**. |
| **Range** | `range` | `flex` | **`all`** | "Take *all* items". |
| **Break** | `break` | `ghosted` | **`cabut`** | "Cabut liao" (Run away / escape). |
| **Continue** | `continue` | `simp` | **`go`** | "Go go go". Keep going. |

### 2.5 Functions & Returns

| Concept | Standard (Go) | CURSED | **Proposed SINGLISH** | Reason / Context |
|:--- |:--- |:--- |:--- |:--- |
| **Return** | `return` | `damn` | **`balek`** | Malay for "return/go back". "Balek kampung". |
| **Defer** | `defer` | `later` | **`nanti`** | Malay for "later". "Nanti do". |

### 2.6 Concurrency (The "Shiok" Part)

| Concept | Standard (Go) | CURSED | **Proposed SINGLISH** | Reason / Context |
|:--- |:--- |:--- |:--- |:--- |
| **Go** | `go` | `stan` | **`chiong`** | To rush/charge. "Cindy *chiong* this task". Implies async speed. |
| **Channel** | `chan` | `dm` | **`lobang`** | A gap/hole/opportunity. You pass things through the *lobang*. |
| **Send** | `<-` | `spill` | **`pass`** | "Pass to lobang". |
| **Receive** | `<-` | (implicit) | **`catch`** | "Catch from lobang". |

### 2.7 Values & Types

| Concept | Standard (Go) | CURSED | **Proposed SINGLISH** | Reason / Context |
|:--- |:--- |:--- |:--- |:--- |
| **True** | `true` | `based` | **`can`** | Positive affirmation. |
| **False** | `false` | `cringe` | **`cannot`** | Negative affirmation. |
| **Nil** | `nil` | `nah` | **`kosong`** | Malay for "empty/zero". |
| **Bool** | `bool` | `lit` | **`bolehtak`** | "Can or not?". The boolean type. |
| **Int** | `int` | `normie` | **`nombor`** | Number. |
| **Int64** | `int64` | `thicc` | **`banyak`** | "Many/Much". Large number. |
| **Float** | `float64` | `meal` | **`point`** | Decimal point. |
| **Complex** | `complex128` | `galaxy_brain` | **`cheem`** | "Cheem" (Profound/Deep). Complex numbers. |
| **String** | `string` | `tea` | **`tar`** | Talk/Speech. |
| **Struct** | `struct` | `squad` | **`barang`** | "Things". A collection of data. |
| **Error** | `error` | `rly` | **`salah`** | "Wrong". Error type. |
| **Panic** | `panic` | `shook` | **`gabra`** | Panic/clumsy confusion. |
| **Pointer** | `*T` | `ඞ` | **`ki`** | Teochew for "point". Points to a value. |
| **Char** | `rune` | `sip` | **`zhi`** | Teochew for Letter/Character. |
| **Recover** | `recover` | `clutch` | **`heng`** | "Heng ah" (Lucky). Recovered from disaster. |

### 2.8 Data Structures & Interfaces

| Concept | Standard (Go) | CURSED | **Proposed SINGLISH** | Reason / Context |
|:--- |:--- |:--- |:--- |:--- |
| **Interface** | `interface` | `vibe_check` (n/a) | **`kaki`** | "Same kaki" (Same group/type). Defines a shared behavior. Use `pattern Foo kaki { ... }` for interface definitions. |
| **Map** | `map` | `tea_spill` | **`menu`** | Key-Value pairs, like a food menu. |
| **Connect** | `concat` | `link` | **`sambung`** | Join strings or lists. |

### 2.9 Built-in Functions & Operators

| Concept | Standard (Go) | CURSED | **Proposed SINGLISH** | Reason / Context |
|:--- |:--- |:--- |:--- |:--- |
| **Make/New** | `make` | `cook` | **`buat`** | Malay for "Make". |
| **Append** | `append` | `thicc` | **`upsize`** | "Upsize the meal" (Make it bigger/add to it). |
| **Delete** | `delete` | `yeet_out` | **`buang`** | Malay for "Throw away". |
| **Len** | `len` | `math` | **`count`** | Simple count. |
| **Close** | `close` | `ghost` | **`kwear`** | Teochew for "Close". |
| **Print** | `print` | `spill` | **`gong`** | Hokkien for "Talk". |
| **And** | `&&` | `and` | **`somemore`** | "This one... somemore that one". |
| **Or** | `\|\|` | `or` | **`or`** | (Standard English is fine). |
| **Not** | `!` | `not` | **`dun`** | "Dun do this". |

## 3. Example Usage

### Hello World

```singlish
kampung main
dapao "fmt"

action boss() {
    gong("Hello Singapore!")
}
```

### FizzBuzz

```singlish
kampung main
dapao "fmt"

action boss() {
    confirm MAX nombor = 100
    got i nombor = 1

    loop i <= MAX {
        nasi i % 15 == 0 {
            gong("FizzBuzz")
        } den nasi i % 3 == 0 {
            gong("Fizz")
        } den nasi i % 5 == 0 {
            gong("Buzz")
        } den {
            gong(i)
        }
        i++
    }
}
```

### Async Worker (Goroutines)

```singlish
kampung main
dapao "fmt"

action worker(ch lobang tar) {
    ch pass "Swee lah!"
    kwear(ch)
}

action boss() {
    got ch = buat(lobang tar, 1)

    // Start goroutine
    chiong worker(ch)

    // Receive from channel
    got result = catch ch
    gong(result)
}
```

### Interface implementation

```singlish
kampung main
dapao "fmt"

pattern Driver kaki {
    Drive() tar
}

pattern Uncle barang {
    Name tar
}

action (u *Uncle) Drive() tar {
    balek "Siam ah! Uncle " + u.Name + " coming!"
}

action boss() {
    got d Driver = &Uncle{Name: "Tan"}
    gong(d.Drive())
}
```

## 4. Summary of Changes

This mapping simplifies the previous mix of Malay/Chinese/English into a cohesive set of "code-switching" terms common in Singlish. It leverages:

- **Action Verbs**: `chiong`, `dapao`, `cabut`.
- **States**: `can`, `cannot`, `kosong`.
- **Structure**: `kampung`, `barang`.
