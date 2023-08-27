package main

import (
	"math/rand"
	"strconv"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Pipe struct {
	posX   int32
	posY   int32
	width  int32
	height int32
	color  rl.Color
}

func main() {
	var screenWidth int32 = 800
	var screenHeight int32 = 450

	var score int
	var didCollide bool

	rl.InitWindow(screenWidth, screenHeight, "Dino")
	rl.SetTargetFPS(60)

	bird_down := rl.LoadImage("assets/bird-down.png")
	bird_up := rl.LoadImage("assets/bird-up.png")

	texture := rl.LoadTextureFromImage(bird_down)

	var x_coords int32 = 20
	var y_coords int32 = screenHeight - 34
	var fixed_y_coords int32 = screenHeight - 34

	rand.New(rand.NewSource(time.Now().UnixNano()))
	var random_height = rand.Intn(200-100+1) + 100
	var random_top_height = rand.Intn(200-100+1) + 100

	Pipes := []Pipe{}
	current_pipe := Pipe{
		posX:   screenWidth - 60,
		posY:   screenHeight - 34,
		width:  30,
		height: int32(random_height),
		color:  rl.Red,
	}

	TopPipes := []Pipe{}
	top_pipe := Pipe{
		posX:   screenWidth - 60,
		posY:   0,
		width:  30,
		height: int32(random_top_height),
		color:  rl.Blue,
	}

	Pipes = append(Pipes, current_pipe)
	TopPipes = append(TopPipes, top_pipe)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.DrawText("Current score: "+strconv.Itoa(score), 0, 0, 30, rl.LightGray)
		rl.DrawTexture(texture, x_coords, y_coords, rl.White)
		rl.ClearBackground(rl.Black)

		if rl.IsKeyDown(rl.KeySpace) {
			texture = rl.LoadTextureFromImage(bird_up)
			if y_coords > 0 {
				y_coords -= 5
			}

		} else {
			texture = rl.LoadTextureFromImage(bird_down)

			if y_coords < fixed_y_coords {
				y_coords += 5
			}
		}

		for io, current_pipe := range Pipes {

			rl.DrawRectangle(current_pipe.posX, screenHeight-current_pipe.height, current_pipe.width, current_pipe.height, current_pipe.color)

			for top_io, top_pipe := range TopPipes {
				rl.DrawRectangle(top_pipe.posX, 0, top_pipe.width, top_pipe.height, top_pipe.color)

				Pipes[io].posX = Pipes[io].posX - 5

				TopPipes[top_io].posX = TopPipes[top_io].posX - 5

				if TopPipes[top_io].posX < 0 && Pipes[io].posX < 0 {
					TopPipes[top_io].posX = int32(rand.Intn(int(screenWidth)-700+1) + 700)
					TopPipes[top_io].height = int32(rand.Intn(200-100+1) + 100)

					Pipes[io].posX = screenWidth
					Pipes[io].height = int32(rand.Intn(200-100+1) + 100)
				}

				bird_rec := rl.NewRectangle(float32(x_coords), float32(y_coords), 34, 24)
				top_pipe_rec := rl.NewRectangle(float32(top_pipe.posX), float32(0), float32(top_pipe.width), float32(top_pipe.height))
				bottom_pipe_rec := rl.NewRectangle(float32(current_pipe.posX), float32(screenHeight-current_pipe.height), float32(current_pipe.width), float32(current_pipe.height))

				if rl.CheckCollisionRecs(bird_rec, top_pipe_rec) || rl.CheckCollisionRecs(bird_rec, bottom_pipe_rec) {
					didCollide = true
					break
				} else {
					score++
				}
			}
		}

		if didCollide {
			rl.UnloadTexture(texture)
			TopPipes = nil
			Pipes = nil
			rl.DrawText("Your final score is: "+strconv.Itoa(score), 30, 40, 30, rl.Red)
		}

		rl.EndDrawing()
	}
	rl.UnloadTexture(texture)
	rl.CloseWindow()
}
