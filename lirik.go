package main

import (
  "fmt"
  "sort"
)

func main() {
  // Mendefinisikan map yang berisi lirik lagu tanpa duplikat
  List_Stanza := map[int]string{
    64: "Stanza I",
    65: "/n",
    2:  "Tanah tumpah darahku,",
    3:  "Jadi pandu ibuku,",
    4:  "Indonesia, tanah airku,",
    5:  "Indonesia, kebangsaanku,",
    6:  "Bangsa dan tanah airku,",
    7:  "Indonesia bersatu,",
    56: "/n",
    35: "Indonesia Raya, merdeka! Merdeka!",
    36: "Hiduplah Indonesia Raya!",
    57: "/n",
    68: "Stanza III",
    69: "/n",
    37: "Indonesia Raya, merdeka! Merdeka!",
    38: "Indonesia, tanah berseri,", // Nilai ini dipilih untuk kunci 38
    39: "Di sanalah aku berdiri,",
    40: "Menjaga ibu sejati,",
    9:  "Hiduplah tanahku,",
    10: "Hiduplah negeriku,",
    26: "Indonesia bahagia!",
    27: "Bangunlah jiwanya,",
    58: "/dn",
    59: "Indonesia Raya, merdeka! Merdeka!",
    60: "/n",
    61: "Stanza II",
    62: "Bangunlah badannya,",
    63: "Indonesia, tanah yang mulia,",
    11: "Sadarilah budinya,",
    31: "Tanah kita yang kaya,",
    54: "Hiduplah Indonesia Raya!",
    49: "Majulah pandunya,",
    50: "Untuk Indonesia Raya.",
    14: "Untuk Indonesia Raya.",
    15: "Di sanalah aku berdiri,",
    16: "Tanahku, negeriku yang kucinta",
    17: "Hiduplah Indonesia Raya!",
    67: "/n",
    19: "Indonesia, tanah pusaka,",
    20: "Pusaka kita semuanya,",
    21: "Marilah kita mendoa,",
    22: "Untuk selama-lamanya,",
    23: "/n",
    24: "Suburlah tanahnya,",
    25: "Suburlah jiwanya,",
    28: "Bangsanya, rakyatnya, semuanya",
    29: "Untuk Indonesia Raya.",
    42: "Tanah yang aku sayangi,", // Nilai ini dipilih untuk kunci 42
    52: "Tanahku, negeriku yang kucinta",
    53: "Marilah kita berjanji,",
    55: "Selamatlah putranya,",
  }

  // Step 1: Mengurutkan kunci dari map
  keys := make([]int, 0, len(List_Stanza))
  for k := range List_Stanza {
    keys = append(keys, k)
  }
  sort.Ints(keys)

  // Step 2: Menampilkan lirik berdasarkan urutan kunci
  for _, k := range keys {
    // Step 3: Cermati penanda "/n" dan "/dn" untuk mencetak baris baru
    if List_Stanza[k] == "/n" || List_Stanza[k] == "/dn" {
      fmt.Println()
    } else {
      fmt.Println(List_Stanza[k])
    }
  }
}
