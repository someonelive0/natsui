# Welcome to Nats-UI Project!

Nats-UI is a GUI project NATS that is a message queue.

Nats-UI use Wails 3 + Vue 3 Typescript, and Vuetify 3.


Hippo icon is from:
https://www.iconarchive.com/show/childrens-book-animals-icons-by-iconarchive/Hippo-icon.html


## License

This software is under MIT license, to visit file [LICENSE](LICENSE) and https://mit-license.org/ for details.


## Getting Started

1. Navigate to your project directory in the terminal.

2. To run your application in development mode, use the following command:

   ```
   wails3 dev

   # frontend dev
   npm run build
   ```

   This will start your application and enable hot-reloading for both frontend and backend changes.

3. To build your application for production, use:

   ```
   wails3 build
   ```

   This will create a production-ready executable in the `build` directory.

   Removing the console window for wails3 projects.
   The quick answer is you can run $env:PRODUCTION="true"; wails3 build.
   For editing pipelines, edit build/windows/Taskfile.yml and customise build.

   ```
   $env:PRODUCTION="true"
   wails3 build
   upx .\bin\natsui.exe
   ```

## Exploring Wails3 Features

Now that you have your project set up, it's time to explore the features that Wails3 offers:

1. **Check out the examples**: The best way to learn is by example. Visit the `examples` directory in the `v3/examples` directory to see various sample applications.

2. **Run an example**: To run any of the examples, navigate to the example's directory and use:

   ```
   go run .
   ```

   Note: Some examples may be under development during the alpha phase.

3. **Explore the documentation**: Visit the [Wails3 documentation](https://v3.wails.io/) for in-depth guides and API references.

4. **Join the community**: Have questions or want to share your progress? Join the [Wails Discord](https://discord.gg/JDdSxwjhGf) or visit the [Wails discussions on GitHub](https://github.com/wailsapp/wails/discussions).

## Project Structure

Take a moment to familiarize yourself with your project structure:

- `frontend/`: Contains your frontend code (HTML, CSS, JavaScript/TypeScript)
- `main.go`: The entry point of your Go backend
- `app.go`: Define your application structure and methods here
- `wails.json`: Configuration file for your Wails project

## Next Steps

1. Modify the frontend in the `frontend/` directory to create your desired UI.
2. Add backend functionality in `main.go`.
3. Use `wails3 dev` to see your changes in real-time.
4. When ready, build your application with `wails3 build`.

Happy coding with Wails3! If you encounter any issues or have questions, don't hesitate to consult the documentation or reach out to the Wails community.
