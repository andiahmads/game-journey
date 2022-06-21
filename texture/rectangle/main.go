package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	maxFrameSpeed = 15
	minFrameSpeed = 1
)

func main() {

	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [texture] example - texture loading and drawing")

	scrafy := rl.LoadTexture("scarfy.png")

	position := rl.NewVector2(350.0, 280.0)
	frameRec := rl.NewRectangle(0, 0, float32(scrafy.Width/2), float32(scrafy.Height))
	currentFrame := float32(0)

	frameCounter := 0
	frameSpeed := 8

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		frameCounter++

		if frameCounter > (60 / frameSpeed) {
			frameCounter = 0
			currentFrame++

			if currentFrame > 5 {
				currentFrame = 0
			}

			frameRec.X = currentFrame * float32(scrafy.Width) / 6
		}

		if rl.IsKeyPressed(rl.KeyRight) {
			frameSpeed++
		} else if rl.IsKeyPressed(rl.KeyLeft) {
			frameSpeed--
		}

		if frameSpeed > maxFrameSpeed {
			frameSpeed = maxFrameSpeed
		} else if frameSpeed < minFrameSpeed {
			frameSpeed = minFrameSpeed
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.DrawTexture(scrafy, 15, 40, rl.White)
		rl.DrawRectangleLines(15, 40, scrafy.Width, scrafy.Height, rl.Lime)
		rl.DrawRectangleLines(15+int32(frameRec.X), 40+int32(frameRec.Y), int32(frameRec.Width), int32(frameRec.Height), rl.Red)
		rl.DrawText("PRESSS RIGHT/LEFT KEYS TO CHANGE SPEED", 290, 240, 10, rl.DarkGray)


		for i := 0; i < maxFrameSpeed; i++ {
			if i < frameSpeed {
				rl.DrawRectangleLines(int32(250+21*i), 205, 20, 20, rl.Red)
			}
			rl.DrawRectangleLines(int32(250+21*i), 205, 20, 20, rl.Maroon)
		}

		rl.DrawTextureRec(scrafy, frameRec, position, rl.White)
		rl.DrawText("(C) scrafy sprite by Eiden Marsal", screenWidth-200, screenHeight-20, 10, rl.Gray)

		rl.EndDrawing()
	}

	rl.UnloadTexture(scrafy)

	rl.CloseWindow()

}
