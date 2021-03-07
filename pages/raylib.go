
package pages

const Raylib = "**raylib** is a cross-platform easy-to-use graphics library, built around\nOpenGL 1.1, 2.1, 3.3 and OpenGL ES 2.0. Even though it is written in C\nit has bindings to over 50 different languages. This tutorial will use C,\nmore specifically C99.\n\n```c\n#include <raylib.h>\n\nint main(void)\n{\n    const int screenWidth = 800;\n    const int screenHeight = 450;\n\n    // Before initialising raylib we can set configuration flags\n    SetConfigFlags(FLAG_MSAA_4X_HINT | FLAG_VSYNC_HINT);\n\n    // raylib doesn't require us to store any instance structures\n    // At the moment raylib can handle only one window at a time\n    InitWindow(screenWidth, screenHeight, \"MyWindow\");\n\n    // Set our game to run at 60 frames-per-second\n    SetTargetFPS(60);\n\n    // Set a key that closes the window\n    // Could be 0 for no key\n    SetExitKey(KEY_DELETE);\n\n    // raylib defines two types of cameras: Camera3D and Camera2D\n    // Camera is a typedef for Camera3D\n    Camera camera = {\n            .position = {0.0f, 0.0f, 0.0f},\n            .target   = {0.0f, 0.0f, 1.0f},\n            .up       = {0.0f, 1.0f, 0.0f},\n            .fovy     = 70.0f,\n            .type     = CAMERA_PERSPECTIVE\n    };\n\n    // raylib supports loading of models, animations, images and sounds\n    // from various different file formats\n    Model myModel = LoadModel(\"my_model.obj\");\n    Font someFont = LoadFont(\"some_font.ttf\");\n\n    // Creates a 100x100 render texture\n    RenderTexture renderTexture = LoadRenderTexture(100, 100);\n\n    // WindowShouldClose checks if the user is closing the window\n    // This might happen using a shortcut, window controls\n    // or the key we set earlier\n    while (!WindowShouldClose())\n    {\n\n        // BeginDrawing needs to be called before any draw call\n        BeginDrawing();\n        {\n\n            // Sets the background to a certain color\n            ClearBackground(BLACK);\n\n            if (IsKeyDown(KEY_SPACE))\n                DrawCircle(400, 400, 30, GREEN);\n\n            // Simple draw text\n            DrawText(\"Congrats! You created your first window!\",\n                     190, // x\n                     200, // y\n                     20,  // font size\n                     LIGHTGRAY\n            );\n\n            // For most functions there are several versions\n            // These are usually postfixed with Ex, Pro, V\n            // or sometimes Rec, Wires (only for 3D), Lines (only for 2D)\n            DrawTextEx(someFont,\n                       \"Text in another font\",\n                       (Vector2) {10, 10},\n                       20, // font size\n                       2,  // spacing\n                       LIGHTGRAY);\n\n            // Required for drawing 3D, has 2D equivalent\n            BeginMode3D(camera);\n            {\n\n                DrawCube((Vector3) {0.0f, 0.0f, 3.0f},\n                         1.0f, 1.0f, 1.0f, RED);\n\n                // White tint when drawing will keep the original color\n                DrawModel(myModel, (Vector3) {0.0f, 0.0f, 3.0f},\n                          1.0f, //Scale\n                          WHITE);\n\n            }\n            // End 3D mode so we can draw normally again\n            EndMode3D();\n\n            // Start drawing onto render texture\n            BeginTextureMode(renderTexture);\n            {\n\n                // It behaves the same as if we just called `BeginDrawing()`\n\n                ClearBackground(RAYWHITE);\n\n                BeginMode3D(camera);\n                {\n\n                    DrawGrid(10, // Slices\n                             1.0f // Spacing\n                    );\n\n                }\n                EndMode3D();\n\n            }\n            EndTextureMode();\n\n            // render textures have a Texture2D field\n            DrawTexture(renderTexture.texture, 40, 378, BLUE);\n\n        }\n        EndDrawing();\n    }\n\n    // Unloading loaded objects\n    UnloadFont(someFont);\n    UnloadModel(myModel);\n\n    // Close window and OpenGL context\n    CloseWindow();\n\n    return 0;\n}\n\n```\n\n## Further reading\nraylib has some [great examples](https://www.raylib.com/examples.html)\nIf you don't like C check out the [raylib bindings](https://github.com/raysan5/raylib/blob/master/BINDINGS.md)"
