### Explanation of the Code:
 * readTasks(): Reads the tasks from the file (taskstxt). If the file doesn’t exist, it returns an empty list.

 * saveTasks(): Saves the tasks back to the file.

 * viewTasks(): Displays the list of tasks to the user.

 * addTask(): Adds a new task to the list and saves it.

 * deleteTask(): Deletes a task by its number (as displayed in the view).

 * main(): Implements a simple CLI menu where users can 
choose to view, add, delete tasks, or exit the program.

   Run the To-Do List CLI Tool
Now you can run the project using:


        go run main.go

### Test the Features

View Tasks: If no tasks exist, it will display "No tasks found." otherwise it will list the tasks with numbers.

Add Task: When prompted, enter the task description. It will be saved in the tasks.txt file.

Delete Task: View the tasks and enter the task number you want to delete.

Exit: Exits the CLI tool.