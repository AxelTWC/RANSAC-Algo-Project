/**
 * Student Name: Axel Tang
 * Student ID: 300164095
 * CSI 2120 Project: Part 1 
 */
import java.util.ArrayList;
import java.util.Iterator;
import java.util.List;
import java.io.File;
import java.io.FileNotFoundException;
import java.util.Scanner;
import java.io.FileWriter;   
import java.io.IOException;  
public class PointCloud {
  List<Point3D> pointcloud = new ArrayList<Point3D>();
  int counter = 0;
    PointCloud(String filename){ //Stores at O(n) <-- Constructor
        try {
          File countLinesFile = new File(filename);
          Scanner countLines = new Scanner(countLinesFile);
          while (countLines.hasNextLine()) {
            countLines.nextLine();
            counter++;
        }
        countLines.close();
          String[][] dataArray = new String[counter][3];
          File pointcloudfile = new File(filename);
          Scanner reader = new Scanner(pointcloudfile);  
          for(int i = 0; i < counter; i++){ //Reading including "x y z"
            if(reader.hasNextLine() && i == 0){ //"x y z"
              String dataXYZSTRING = reader.nextLine();
            }
            if(reader.hasNextLine() && i != 0){ //other points 
              String data = reader.nextLine();
              String[] temp = data.split("\\s+");
              for(int index = 0; index < 3; index++){
                dataArray[i-1][index] = temp[index];
              }
            }
          }
          for(int test = 0; test < counter-1; test ++){ //Up size - 1 counting for xyz
            Point3D point3dObj = new Point3D(Double.parseDouble(dataArray[test][0]), Double.parseDouble(dataArray[test][1]), Double.parseDouble(dataArray[test][2]));
            pointcloud.add(point3dObj);
          }
          reader.close();
        } catch (FileNotFoundException e) {
          System.out.println("An error occurred.");
          e.printStackTrace();
        }
  }
    PointCloud(){} //Empty Constructor
    public void addPoint(Point3D pt){ //Add a point to point cloud
      pointcloud.add(pt);
    }

    Point3D getPoint(){ //Get Random Point
      int randN = (int) ((Math.random() * (counter-1 - 0)) + 0);
      return pointcloud.get(randN);
    }
    public int getSize(){ //get size of point cloud
      return counter;
    }
    public void save(String filename){ //Write PointCloud File (.xyz) format
      try {
        FileWriter writer = new FileWriter(filename);
        writer.write("x y z");
        writer.write("\n");
        for(int i = 0; i < pointcloud.size(); i++){
          writer.write(pointcloud.get(i).getX() + " "+  pointcloud.get(i).getY() + " " + pointcloud.get(i).getZ());
          writer.write("\n");
        }
        writer.close();
        System.out.println("Successfully Wrote PointCloud into " + filename);
      } catch (IOException e) {
        System.out.println("An error occurred.");
        e.printStackTrace();
      }
    }
    
    Iterator<Point3D> iterator(){
      Iterator<Point3D> iteratorPointCloud = pointcloud.iterator();
      return iteratorPointCloud;
    }

    // public static void main(String[] args){ //Test Function
    //     long start,end;
    //     double tim;
    //     start=System.currentTimeMillis();
    //     PointCloud tempCloud = new PointCloud();
    //     tempCloud = new PointCloud("PointCloud1.xyz");
    //     Point3D test = new Point3D(0, 0, 0);
    //     tempCloud.addPoint(test);
    //     tempCloud.save("Test.xyz");
    //     end=System.currentTimeMillis();
    //     tim=(end-start)/1000.0;
    //     System.out.println("Timer For Creating File: ");
    //     System.out.println(tim + "seconds");
    //     // Point3D forTest = tempCloud.getPoint();
    //     // System.out.println(forTest.getX() + " " + forTest.getY() + " " + forTest.getZ());
    // }
}

