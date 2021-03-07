
package pages

const Opengl = "**Open Graphics Library** (**OpenGL**) is a cross-language cross-platform application programming interface\n(API) for rendering 2D computer graphics and 3D vector graphics.<sup>[1]</sup> In this tutorial we will be\nfocusing on modern OpenGL from 3.3 and above, ignoring \"immediate-mode\", Displaylists and\nVBO's without use of Shaders.\nI will be using C++ with SFML for window, image and context creation aswell as GLEW\nfor modern OpenGL extensions, though there are many other librarys available.\n\n```cpp\n// Creating an SFML window and OpenGL basic setup.\n#include <GL/glew.h>\n#include <GL/gl.h>\n#include <SFML/Graphics.h>\n#include <iostream>\n\nint main() {\n    // First we tell SFML how to setup our OpenGL context.\n    sf::ContextSettings context{ 24,   // depth buffer bits\n                                  8,   // stencil buffer bits\n                                  4,   // MSAA samples\n                                  3,   // major opengl version\n                                  3 }; // minor opengl version\n    // Now we create the window, enable VSync\n    // and set the window active for OpenGL.\n    sf::Window window{ sf::VideoMode{ 1024, 768 },\n                       \"opengl window\",\n                       sf::Style::Default,\n\t\t       context };\n    window.setVerticalSyncEnabled(true);\n    window.setActive(true);\n    // After that we initialise GLEW and check if an error occured.\n    GLenum error;\n    glewExperimental = GL_TRUE;\n    if ((err = glewInit()) != GLEW_OK)\n        std::cout << glewGetErrorString(err) << std::endl;\n    // Here we set the color glClear will clear the buffers with.\n    glClearColor(0.0f,    // red\n                 0.0f,    // green\n                 0.0f,    // blue\n                 1.0f);   // alpha\n    // Now we can start the event loop, poll for events and draw objects.\n    sf::Event event{ };\n    while (window.isOpen()) {\n        while (window.pollEvent(event)) {\n            if (event.type == sf::Event::Closed)\n                window.close;\n        }\n        // Tell OpenGL to clear the color buffer\n        // and the depth buffer, this will clear our window.\n        glClear(GL_COLOR_BUFFER_BIT | GL_DEPTH_BUFFER_BIT);\n        // Flip front- and backbuffer.\n        window.display();\n    }\n    return 0;\n}\n```\n\n## Loading Shaders\n\nAfter creating a window and our event loop we should create a function,\nthat sets up our shader program.\n\n```cpp\nGLuint createShaderProgram(const std::string& vertexShaderPath,\n                           const std::string& fragmentShaderPath) {\n    // Load the vertex shader source.\n    std::stringstream ss{ };\n    std::string vertexShaderSource{ };\n    std::string fragmentShaderSource{ };\n    std::ifstream file{ vertexShaderPath };\n    if (file.is_open()) {\n        ss << file.rdbuf();\n        vertexShaderSource = ss.str();\n        file.close();\n    }\n    // Clear the stringstream and load the fragment shader source.\n    ss.str(std::string{ });\n    file.open(fragmentShaderPath);\n    if (file.is_open()) {\n        ss << file.rdbuf();\n        fragmentShaderSource = ss.str();\n        file.close();\n    }\n    // Create the program.\n    GLuint program = glCreateProgram();\n    // Create the shaders.\n    GLuint vertexShader = glCreateShader(GL_VERTEX_SHADER);\n    GLuint fragmentShader = glCreateShader(GL_FRAGMENT_SHADER);\n    // Now we can load the shader source into the shader objects and compile them.\n    // Because glShaderSource() wants a const char* const*,\n    // we must first create a const char* and then pass the reference.\n    const char* cVertexSource = vertexShaderSource.c_str();\n    glShaderSource(vertexShader,     // shader\n                   1,                // number of strings\n                   &cVertexSource,   // strings\n                   nullptr);         // length of strings (nullptr for 1)\n    glCompileShader(vertexShader);\n    // Now we have to do the same for the fragment shader.\n    const char* cFragmentSource = fragmentShaderSource.c_str();\n    glShaderSource(fragmentShader, 1, &cFragmentSource, nullptr);\n    glCompileShader(fragmentShader);\n    // After attaching the source and compiling the shaders,\n    // we attach them to the program;\n    glAttachShader(program, vertexShader);\n    glAttachShader(program, fragmentShader);\n    glLinkProgram(program);\n    // After linking the shaders we should detach and delete\n    // them to prevent memory leak.\n    glDetachShader(program, vertexShader);\n    glDetachShader(program, fragmentShader);\n    glDeleteShader(vertexShader);\n    glDeleteShader(fragmentShader);\n    // With everything done we can return the completed program.\n    return program;\n}\n```\n\nIf you want to check the compilation log you can add the following between <code>glCompileShader()</code> and <code>glAttachShader()</code>.\n\n```cpp\nGLint logSize = 0;\nstd::vector<GLchar> logText{ };\nglGetShaderiv(vertexShader,         // shader\n              GL_INFO_LOG_LENGTH,   // requested parameter\n              &logSize);            // return object\nif (logSize > 0) {\n    logText.resize(logSize);\n    glGetShaderInfoLog(vertexShader,      // shader\n                       logSize,           // buffer length\n                       &logSize,          // returned length\n                       logText.data());   // buffer\n    std::cout << logText.data() << std::endl;\n}\n```\n\nThe same is possibile after <code>glLinkProgram()</code>, just replace <code>glGetShaderiv()</code> with <code>glGetProgramiv()</code>\nand <code>glGetShaderInfoLog()</code> with <code>glGetProgramInfoLog()</code>.\n\n```cpp\n// Now we can create a shader program with a vertex and a fragment shader.\n// ...\nglClearColor(0.0f, 0.0f, 0.0f, 1.0f);\n\nGLuint program = createShaderProgram(\"vertex.glsl\", \"fragment.glsl\");\n\nsf::Event event{ };\n// ...\n// We also have to delete the program at the end of the application.\n// ...\n    }\n    glDeleteProgram(program);\t\n    return 0;\n}\n// ...\n```\n\nOfcourse we have to create the vertex and fragment shader before we can load them,\nso lets create two basic shaders.\n\n**Vertex Shader**\n\n```glsl\n// Declare which version of GLSL we use.\n// Here we declare, that we want to use the OpenGL 3.3 version of GLSL.\n#version 330 core\n// At attribute location 0 we want an input variable of type vec3,\n// that contains the position of the vertex.\n// Setting the location is optional, if you don't set it you can ask for the\n// location with glGetAttribLocation().\nlayout(location = 0) in vec3 position;\n// Every shader starts in it's main function.\nvoid main() {\n    // gl_Position is a predefined variable that holds\n    // the final vertex position.\n    // It consists of a x, y, z and w coordinate.\n    gl_Position = vec4(position, 1.0);\n}\n```\n\n**Fragment Shader**\n\n```glsl\n#version 330 core\n// The fragment shader does not have a predefined variable for\n// the vertex color, so we have to define a output vec4,\n// that holds the final vertex color.\nout vec4 outColor;\n\nvoid main() {\n    // We simply set the ouput color to red.\n    // The parameters are red, green, blue and alpha.\n    outColor = vec4(1.0, 0.0, 0.0, 1.0);\n}\n```\n\n## VAO and VBO\nNow we need to define some vertex position we can pass to our shaders. Lets define a simple 2D quad.\n\n```cpp\n// The vertex data is defined in a counter-clockwise way,\n// as this is the default front face.\nstd::vector<float> vertexData {\n    -0.5f,  0.5f, 0.0f,\n    -0.5f, -0.5f, 0.0f,\n     0.5f, -0.5f, 0.0f,\n     0.5f,  0.5f, 0.0f\n};\n// If you want to use a clockwise definition, you can simply call\nglFrontFace(GL_CW);\n// Next we need to define a Vertex Array Object (VAO).\n// The VAO stores the current state while its active.\nGLuint vao = 0;\nglGenVertexArrays(1, &vao);\nglBindVertexArray(vao);\n// With the VAO active we can now create a Vertex Buffer Object (VBO).\n// The VBO stores our vertex data.\nGLuint vbo = 0;\nglGenBuffers(1, &vbo);\nglBindBuffer(GL_ARRAY_BUFFER, vbo);\n// For reading and copying there are also GL_*_READ and GL_*_COPY,\n// if your data changes more often use GL_DYNAMIC_* or GL_STREAM_*.\nglBufferData(GL_ARRAY_BUFFER,     // target buffer\n             sizeof(vertexData[0]) * vertexData.size(),   // size\n             vertexData.data(),   // data\n             GL_STATIC_DRAW);     // usage\n// After filling the VBO link it to the location 0 in our vertex shader,\n// which holds the vertex position.\n// ...\n// To ask for the attibute location, if you haven't set it:\nGLint posLocation = glGetAttribLocation(program, \"position\");\n// ..\nglEnableVertexAttribArray(0);\nglVertexAttribPointer(0, 3,       // location and size\n                      GL_FLOAT,   // type of data\n                      GL_FALSE,   // normalized (always false for floats)\n                      0,          // stride (interleaved arrays)\n                      nullptr);   // offset (interleaved arrays)\n// Everything should now be saved in our VAO and we can unbind it and the VBO.\nglBindVertexArray(0);\nglBindBuffer(GL_ARRAY_BUFFER, 0);\n// Now we can draw the vertex data in our render loop.\n// ...\nglClear(GL_COLOR_BUFFER_BIT);\n// Tell OpenGL we want to use our shader program.\nglUseProgram(program);\n// Binding the VAO loads the data we need.\nglBindVertexArray(vao);\n// We want to draw a quad starting at index 0 of the VBO using 4 indices.\nglDrawArrays(GL_QUADS, 0, 4);\nglBindVertexArray(0);\nwindow.display();\n// ...\n// Ofcource we have to delete the allocated memory for the VAO and VBO at\n// the end of our application.\n// ...\nglDeleteBuffers(1, &vbo);\nglDeleteVertexArrays(1, &vao);\nglDeleteProgram(program);\nreturn 0;\n// ...\n```\n\nYou can find the current code here: [OpenGL - 1](https://pastebin.com/W8jdmVHD).\n\n## More VBO's and Color\nLet's create another VBO for some colors.\n\n```cpp\nstd::vector<float> colorData {\n    1.0f, 0.0f, 0.0f,\n    0.0f, 1.0f, 0.0f,\n    0.0f, 0.0f, 1.0f,\n    1.0f, 1.0f, 0.0f\n};\n```\n\nNext we can simply change some previous parameters to create a second VBO for our colors.\n\n```cpp\n// ...\nGLuint vbo[2];\nglGenBuffers(2, vbo);\nglBindBuffer(GL_ARRAY_BUFFER, vbo[0]);\n// ...\nglDeleteBuffers(2, vbo);\n/ ...\n// With these changes made we now have to load our color data into the new VBO\n// ...\nglVertexAttribPointer(0, 3, GL_FLOAT, GL_FALSE, 0, nullptr);\n\nglBindBuffer(GL_ARRAY_BUFFER, vbo[1]);\nglBufferData(GL_ARRAY_BUFFER, sizeof(colorData[0]) * colorData.size(),\n             colorData.data(), GL_STATIC_DRAW);\nglEnableVertexAttribArray(1);\nglVertexAttribPointer(1, 3, GL_FLOAT, GL_FALSE, 0, nullptr);\n\nglBindVertexArray(0);  \n// ...\n```\n\nNext we have to change our vertex shader to pass the color data to the fragment shader.<br>\n**Vertex Shader**\n\n```glsl\n#version 330 core\n\nlayout(location = 0) in vec3 position;\n// The new location has to differ from any other input variable.\n// It is the same index we need to pass to\n// glEnableVertexAttribArray() and glVertexAttribPointer().\nlayout(location = 1) in vec3 color;\n\nout vec3 fColor;\n\nvoid main() {\n    fColor = color;\n    gl_Position = vec4(position, 1.0);\n}\n```\n\n**Fragment Shader**\n\n```glsl\n#version 330 core\n\nin vec3 fColor;\n\nout vec4 outColor;\n\nvoid main() {\n    outColor = vec4(fColor, 1.0);\n}\n```\n\nWe define a new input variable ```color``` which represents our color data, this data\nis passed on to ```fColor```, which is an output variable of our vertex shader and\nbecomes an input variable for our fragment shader.\nIt is imporatant that variables passed between shaders have the exact same name\nand type.\n\n## Handling VBO's\n\n```cpp\n// If you want to completely clear and refill a VBO use glBufferData(),\n// just like we did before.\n// ...\n// There are two mains ways to update a subset of a VBO's data.\n// To update a VBO with existing data\nstd::vector<float> newSubData {\n\t-0.25f, 0.5f, 0.0f\n};\nglBindBuffer(GL_ARRAY_BUFFER, vbo[0]);\nglBufferSubData(GL_ARRAY_BUFFER,      // target buffer\n                0,                    // offset\n                sizeof(newSubData[0]) * newSubData.size(),   // size\n                newSubData.data());   // data\n// This would update the first three values in our vbo[0] buffer.\n// If you want to update starting at a specific location just set the second\n// parameter to that value and multiply by the types size.\n// ...\n// If you are streaming data, for example from a file,\n// it is faster to directly pass the data to the buffer.\n// Other access values are GL_READ_ONLY and GL_READ_WRITE.\nglBindBuffer(GL_ARRAY_BUFFER, vbo[0]);\n// You can static_cast<float*>() the void* to be more safe.\nvoid* Ptr = glMapBuffer(GL_ARRAY_BUFFER,   // buffer to map\n                        GL_WRITE_ONLY);    // access to buffer\nmemcpy(Ptr, newSubData.data(), sizeof(newSubData[0]) * newSubData.size());\n// To copy to a specific location add a destination offset to memcpy().\nglUnmapBuffer(GL_ARRAY_BUFFER);\n// ...\n// There is also a way to copy data from one buffer to another,\n// If we have two VBO's vbo[0] and vbo[1], we can copy like so\n// You can also read from GL_ARRAY_BUFFER.\nglBindBuffer(GL_COPY_READ_BUFFER, vbo[0]);\n// GL_COPY_READ_BUFFER and GL_COPY_WRITE_BUFFER are specifically for\n// copying buffer data.\nglBindBuffer(GL_COPY_WRITE_BUFFER, vbo[1]);\nglCopyBufferSubData(GL_COPY_READ_BUFFER,    // read buffer\n                    GL_COPY_WRITE_BUFFER,   // write buffer\n                    0, 0,                   // read and write offset\n                    sizeof(vbo[0]) * 3);    // copy size\n// This will copy the first three elements from vbo[0] to vbo[1].\n```\n\n## Uniforms\n\n**Fragment Shader**\n\n```glsl\n// Uniforms are variables like in and out, however,\n// we can change them easily by passing new values with glUniform().\n// Lets define a time variable in our fragment shader.\n#version 330 core\n// Unlike a in/out variable we can use a uniform in every shader,\n// without the need to pass it to the next one, they are global.\n// Don't use locations already used for attributes!\n// Uniform layout locations require OpenGL 4.3!\nlayout(location = 10) uniform float time;\n\nin vec3 fColor;\n\nout vec4 outColor;\n\nvoid main() {\n    // Create a sine wave from 0 to 1 based on the time passed to the shader.\n    float factor = (sin(time * 2) + 1) / 2;\n    outColor = vec4(fColor.r * factor, fColor.g * factor, fColor.b * factor, 1.0);\n}\n```\n\nBack to our source code.\n\n```cpp\n// If we haven't set the layout location, we can ask for it.\nGLint timeLocation = glGetUniformLocation(program, \"time\");\n// ...\n// Also we should define a Timer counting the current time.\nsf::Clock clock{ };\n// In out render loop we can now update the uniform every frame.\n    // ...\n    window.display();\n    glUniform1f(10,   // location\n                clock.getElapsedTime().asSeconds());   // data\n}\n// ...\n```\n\nWith the time getting updated every frame the quad should now be changing from\nfully colored to pitch black.\nThere are different types of glUniform() you can find simple documentation here:\n[glUniform - OpenGL Refpage](https://www.khronos.org/registry/OpenGL-Refpages/gl4/html/glUniform.xhtml)\n\n## Indexing and IBO's\n\nElement Array Buffers or more commonly Index Buffer Objects (IBO) allow us to use the\nsame vertex data again which makes drawing a lot easier and faster. here's an example:\n\n```cpp\n// Lets create a quad from two rectangles.\n// We can simply use the old vertex data from before.\n// First, we have to create the IBO.\n// The index is referring to the first declaration in the VBO.\nstd::vector<unsigned int> iboData {\n    0, 1, 2,\n    0, 2, 3\n};\n// That's it, as you can see we could reuse 0 - the top left\n// and 2 - the bottom right.\n// Now that we have our data, we have to fill it into a buffer.\n// Note that this has to happen between the two glBindVertexArray() calls,\n// so it gets saved into the VAO.\nGLuint ibo = 0;\nglGenBufferrs(1, &ibo);\nglBindBuffer(GL_ELEMENT_ARRAY_BUFFER, ibo);\nglBufferData(GL_ELEMENT_ARRAY_BUFFER, sizeof(iboData[0]) * iboData.size(),\n             iboData.data(), GL_STATIC_DRAW);\n// Next in our render loop, we replace glDrawArrays() with:\nglDrawElements(GL_TRIANGLES, iboData.size(), GL_UNSINGED_INT, nullptr);\n// Remember to delete the allocated memory for the IBO.\n```\n\nYou can find the current code here: [OpenGL - 2](https://pastebin.com/R3Z9ACDE).\n\n## Textures\n\nTo load out texture we first need a library that loads the data, for simplicity I will be\nusing SFML, however there are a lot of librarys for loading image data.\n\n```cpp\n// Lets save we have a texture called \"my_tex.tga\", we can load it with:\nsf::Image image;\nimage.loadFromFile(\"my_tex.tga\");\n// We have to flip the texture around the y-Axis, because OpenGL's texture\n// origin is the bottom left corner, not the top left.\nimage.flipVertically();\n// After loading it we have to create a OpenGL texture.\nGLuint texture = 0;\nglGenTextures(1, &texture);\nglBindTexture(GL_TEXTURE_2D, texture);\n// Specify what happens when the coordinates are out of range.\nglTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_WRAP_S, GL_REPEAT);\nglTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_WRAP_T, GL_REPEAT);\n// Specify the filtering if the object is very large.\nglTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MIN_FILTER, GL_LINEAR);\nglTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MAG_FILTER, GL_LINEAR);\n// Load the image data to the texture.\nglTexImage2D(GL_TEXTURE_2D, 0, GL_RGBA, image.getSize().x, image.getSize().y,\n             0, GL_RGBA, GL_UNSIGNED_BYTE, image.getPixelsPtr());\n// Unbind the texture to prevent modifications.\nglBindTexture(GL_TEXTURE_2D, 0);\n// Delete the texture at the end of the application.\n// ...\nglDeleteTextures(1, &texture);\n```\n\nOfcourse there are more texture formats than only 2D textures,\nYou can find further information on parameters here:\n[glBindTexture - OpenGL Refpage](https://www.khronos.org/registry/OpenGL-Refpages/gl4/html/glBindTexture.xhtml)<br>\n[glTexImage2D - OpenGL Refpage](https://www.khronos.org/registry/OpenGL-Refpages/gl4/html/glTexImage2D.xhtml)<br>\n[glTexParameter - OpenGL Refpage](https://www.khronos.org/registry/OpenGL-Refpages/gl4/html/glTexParameter.xhtml)<br>\n\n```cpp\n// With the texture created, we now have to specify the UV,\n// or in OpenGL terms ST coordinates.\nstd::vector<float> texCoords {\n    // The texture coordinates have to match the triangles/quad\n    // definition.\n    0.0f, 1.0f,\t   // start at top-left\n    0.0f, 0.0f,\t   // go round counter-clockwise\n    1.0f, 0.0f,\n    1.0f, 1.0f     // end at top-right\n};\n// Now we increase the VBO's size again just like we did for the colors.\n// ...\nGLuint vbo[3];\nglGenBuffers(3, vbo);\n// ...\nglDeleteBuffers(3, vbo);\n// ...\n// Load the texture coordinates into the new buffer.\nglBindBuffer(GL_ARRAY_BUFFER, vbo[2]);\nglBufferData(GL_ARRAY_BUFFER, sizeof(texCoords[0]) * texCoords.size(),\n             texCoords.data(), GL_STATIC_DRAW);\nglEnableVertexAttribArray(2);\nglVertexAttribPointer(2, 2, GL_FLOAT, GL_FALSE, 0, nullptr);\n// Because the VAO does not store the texture we have to bind it before drawing.\n// ...\nglBindVertexArray(vao);\nglBindTexture(GL_TEXTURE_2D, texture);\nglDrawElements(GL_TRIANGLES, iboData.size(), GL_UNSINGED_INT, nullptr);\n// ...\n```\n\nChange the shaders to pass the data to the fragment shader.<br>\n\n**Vertex Shader**\n\n```glsl\n#version 330 core\n\nlayout(location = 0) in vec3 position;\nlayout(location = 1) in vec3 color;\nlayout(location = 2) in vec2 texCoords;\n\nout vec3 fColor;\nout vec2 fTexCoords;\n\nvoid main() {\n    fColor = color;\n    fTexCoords = texCoords;\n    gl_Position = vec4(position, 1.0);\n}\n```\n\n**Fragment Shader**\n\n```glsl\n#version 330 core\n// sampler2D represents our 2D texture.\nuniform sampler2D tex;\nuniform float time;\n\nin vec3 fColor;\nin vec2 fTexCoords;\n\nout vec4 outColor;\n\nvoid main() {\n    // texture() loads the current texure data at the specified texture coords,\n    // then we can simply multiply them by our color.\n    outColor = texture(tex, fTexCoords) * vec4(fColor, 1.0);\n}\n```\n\nYou can find the current code here: [OpenGL - 3](https://pastebin.com/u3bcwM6q)\n\n## Matrix Transformation\n\n**Vertex Shader**\n\n```glsl\n#version 330 core\n\nlayout(location = 0) in vec3 position;\nlayout(location = 1) in vec3 color;\nlayout(location = 2) in vec2 texCoords;\n// Create 2 4x4 matricies, 1 for the projection matrix\n// and 1 for the model matrix.\n// Because we draw in a static scene, we don't need a view matrix.\nuniform mat4 projection;\nuniform mat4 model;\n\nout vec3 fColor;\nout vec2 fTexCoords;\n\nvoid main() {\n    fColor = color;\n    fTexCoords = texCoords;\n    // Multiplay the position by the model matrix and then by the\n    // projection matrix.\n    // Beware order of multiplication for matricies!\n    gl_Position = projection * model * vec4(position, 1.0);\n}\n```\n\nIn our source we now need to change the vertex data, create a model- and a projection matrix.\n\n```cpp\n// The new vertex data, counter-clockwise declaration.\nstd::vector<float> vertexData {  \n    0.0f, 1.0f, 0.0f,   // top left\n    0.0f, 0.0f, 0.0f,   // bottom left\n    1.0f, 0.0f, 0.0f,   // bottom right\n    1.0f, 1.0f, 0.0f    // top right\n};\n// Request the location of our matricies.\nGLint projectionLocation = glGetUniformLocation(program, \"projection\");\nGLint modelLocation = glGetUniformLocation(program, \"model\");\n// Declaring the matricies.\n// Orthogonal matrix for a 1024x768 window.\nstd::vector<float> projection {  \n    0.001953f,       0.0f,  0.0f, 0.0f,\n         0.0f, -0.002604f,  0.0f, 0.0f,\n         0.0f,       0.0f, -1.0f, 0.0f,\n        -1.0f,       1.0f,  0.0f, 1.0f\n};\n// Model matrix translating to x 50, y 50\n// and scaling to x 200, y 200.\nstd::vector<float> model {  \n    200.0f,   0.0f, 0.0f, 0.0f,\n      0.0f, 200.0f, 0.0f, 0.0f,\n      0.0f,   0.0f, 1.0f, 0.0f,\n     50.0f,  50.0f, 0.0f, 1.0f\n};\n// Now we can send our calculated matricies to the program.\nglUseProgram(program);\nglUniformMatrix4fv(projectionLocation,   // location\n                   1,                    // count\n                   GL_FALSE,             // transpose the matrix\n                   projection.data());   // data\nglUniformMatrix4fv(modelLocation, 1, GL_FALSE, model.data());\nglUseProgram(0);\n// The glUniform*() calls have to be done, while the program is bound.\n```\n\nThe application should now display the texture at the defined position and size.<br>\nYou can find the current code here: [OpenGL - 4](https://pastebin.com/9ahpFLkY)\n\n```cpp\n// There are many math librarys for OpenGL, which create\n// matricies and vectors, the most used in C++ is glm (OpenGL Mathematics).\n// Its a header only library.\n// The same code using glm would look like:\nglm::mat4 projection{ glm::ortho(0.0f, 1024.0f, 768.0f, 0.0f) };\nglUniformMatrix4fv(projectionLocation, 1, GL_FALSE,\n                   glm::value_ptr(projection));\n// Initialise the model matrix to the identity matrix, otherwise every\n// multiplication would be 0.\nglm::mat4 model{ 1.0f };\nmodel = glm::translate(model, glm::vec3{ 50.0f, 50.0f, 0.0f });\nmodel = glm::scale(model, glm::vec3{ 200.0f, 200.0f, 0.0f });\nglUniformMatrix4fv(modelLocation, 1, GL_FALSE,\n                   glm::value_ptr(model));\n```\n\n## Geometry Shader\n\nGemoetry shaders were introduced in OpenGL 3.2, they can produce vertices\nthat are send to the rasterizer. They can also change the primitive type e.g.\nthey can take a point as an input and output other primitives.\nGeometry shaders are inbetween the vertex and the fragment shader.\n\n**Vertex Shader**\n\n```glsl\n#version 330 core\n\nlayout(location = 0) in vec3 position;\nlayout(location = 1) in vec3 color;\n// Create an output interface block passed to the next shadaer stage.\n// Interface blocks can be used to structure data passed between shaders.\nout VS_OUT {\n    vec3 color;\n} vs_out;\n\nvoid main() {\n    vs_out.color = color\n    gl_Position = vec4(position, 1.0);\n}\n```\n\n**Geometry Shader**\n\n```glsl\n#version 330 core\n// The geometry shader takes in points.\nlayout(points) in;\n// It outputs a triangle every 3 vertices emitted.\nlayout(triangle_strip, max_vertices = 3) out;\n// VS_OUT becomes an input variable in the geometry shader.\n// Every input to the geometry shader in treated as an array.\nin VS_OUT {\n    vec3 color;\n} gs_in[];\n// Output color for the fragment shader.\n// You can also simply define color as 'out vec3 color',\n// If you don't want to use interface blocks.\nout GS_OUT {\n    vec3 color;\n} gs_out;\n\nvoid main() {\n    // Each emit calls the fragment shader, so we set a color for each vertex.\n    gs_out.color = mix(gs_in[0].color, vec3(1.0, 0.0, 0.0), 0.5);\n    // Move 0.5 units to the left and emit the new vertex.\n    // gl_in[] is the current vertex from the vertex shader, here we only\n    // use 0, because we are receiving points.\n    gl_Position = gl_in[0].gl_Position + vec4(-0.5, 0.0, 0.0, 0.0);\n    EmitVertex();\n    gs_out.color = mix(gs_in[0].color, vec3(0.0, 1.0, 0.0), 0.5);\n    // Move 0.5 units to the right and emit the new vertex.\n    gl_Position = gl_in[0].gl_Position + vec4(0.5, 0.0, 0.0, 0.0);\n    EmitVertex();\n    gs_out.color = mix(gs_in[0].color, vec3(0.0, 0.0, 1.0), 0.5);\n    // Move 0.5 units up and emit the new vertex.\n    gl_Position = gl_in[0].gl_Position + vec4(0.0, 0.75, 0.0, 0.0);\n    EmitVertex();\n    EndPrimitive();\n}\n```\n\n**Fragment Shader**\n\n```glsl\nin GS_OUT {\n    vec3 color;\n} fs_in;\n\nout vec4 outColor;\n\nvoid main() {\n    outColor = vec4(fs_in.color, 1.0);\n}\n```\n\nIf you now store a single point with a single color in a VBO and draw them,\nyou should see a triangle, with your color mixed half way between\nred, green and blue on each vertex.\n\n\n## Quotes\n<sup>[1]</sup>[OpenGL - Wikipedia](https://en.wikipedia.org/wiki/OpenGL)\n\n## Books\n\n- OpenGL Superbible - Fifth Edition (covering OpenGL 3.3)\n- OpenGL Programming Guide - Eighth Edition (covering OpenGL 4.3)"