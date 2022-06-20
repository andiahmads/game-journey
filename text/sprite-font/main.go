package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {

	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "Raylib [text] example- sprite fonts usage")

	msg1 := "THIS A custom SPRITE FONT"
	msg2 := "...and this is ANOTHER CUSTOEM FONT"
	msg3 := "... and a THIRD one! GREAT! :D"

	font1 := rl.LoadFont("font/custom_mecha.png")
	font2 := rl.LoadFont("font/custom_alagard.png")
	font3 := rl.LoadFont("font/custom_jupiter_crash.png")

	var fontPosition1, fontPosition2, fontPosition3 rl.Vector2

	fontPosition1.X = float32(screenWidth)/2 - rl.MeasureTextEx(font1, msg1, float32(font1.BaseSize), -3).X/2
	fontPosition1.Y = float32(screenHeight)/2 - float32(font1.BaseSize)/2 - 80

	fontPosition2.X = float32(screenWidth)/2 - rl.MeasureTextEx(font2, msg2, float32(font2.BaseSize), -2).X/2
	fontPosition2.Y = float32(screenHeight)/2 - float32(font2.BaseSize)/2 - 10

	fontPosition3.X = float32(screenWidth)/2 - rl.MeasureTextEx(font3, msg3, float32(font3.BaseSize), 2).X/2
	fontPosition3.Y = float32(screenHeight)/2 - float32(font3.BaseSize)/2 + 50

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.DrawTextEx(font1, msg1, fontPosition1, float32(font1.BaseSize), -3, rl.White)
		rl.DrawTextEx(font2, msg2, fontPosition2, float32(font2.BaseSize), -2, rl.White)
		rl.DrawTextEx(font3, msg3, fontPosition3, float32(font3.BaseSize), 2, rl.White)

		rl.EndDrawing()
	}

	rl.UnloadFont(font1)
	rl.UnloadFont(font2)
	rl.UnloadFont(font3)

	rl.CloseWindow()

}
