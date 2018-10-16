/*
Copyright (c) 2018 Tomasz "VedVid" Nowakowski

This software is provided 'as-is', without any express or implied
warranty. In no event will the authors be held liable for any damages
arising from the use of this software.

Permission is granted to anyone to use this software for any purpose,
including commercial applications, and to alter it and redistribute it
freely, subject to the following restrictions:

1. The origin of this software must not be misrepresented; you must not
   claim that you wrote the original software. If you use this software
   in a product, an acknowledgment in the product documentation would be
   appreciated but is not required.
2. Altered source versions must be plainly marked as such, and must not be
   misrepresented as being the original software.
3. This notice may not be removed or altered from any source distribution.
*/

package main

import blt "bearlibterminal"
import "fmt"

func main() {
	player, err := NewPlayer(PlayerLayer, 1, 1, "@", "white", "white", true, true, false)
	if err != nil {
		fmt.Println(err)
	}
	enemy, err := NewCreature(MonstersLayer, 10, 10, "T", "green", "green", false, true, false)
	if err != nil {
		fmt.Println(err)
	}
	var actors = Monsters{player, enemy}
	var objs = Objects{}
	cells := InitializeEmptyMap()
	for {
		RenderAll(cells, objs, actors)
		key := blt.Read()
		if key == blt.TK_ESCAPE {
			break
		} else {
			Controls(key, player)
			if enemy.DistanceTo(player.X, player.Y) > 1 {
				enemy.MoveTowards(cells, player.X, player.Y)
			}
		}
	}
	blt.Close()
}

func init() {
	InitializeFOVTables()
	InitializeBLT()
}
