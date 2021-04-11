# Maze Runner

Maze Runner is a console tool for generating, solving and visualizing mazes.

The randomized Kruskal's algorithm is used to generate the maze, and the A* algorithm is used to find the path from the entrance to the exit of the maze. The maze can be saved and loaded in text form.

It is written in pure Java and has a detailed javadoc.

![](maze.gif)

## Running

It is a simple Java application and it only requires JDK 11 on your machine.

**1. Clone the application**

```bash
git clone https://github.com/Chatrz/DSA/Maze Project/
cd maze-runner
```

**2. Compile**

```bash
javac -sourcepath ./src -d bin src/maze/Main.java
```

**2. Run**

```bash
java -classpath ./bin maze.Main
```

**Note:** you may need to adjust the line space in order to have a continuous maze in the console.
