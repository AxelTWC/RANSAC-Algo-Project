/**
 * Student Name: Axel Tang
 * Student ID: 300164095
 * Student Email: wtang102@uottawa.ca
 * CSI 2120 Project: Part 1 
 */

//Doing Approach 2: Calculating the three most dominant planes directly from the original point cloud, without removing any points.
// (Please Contact me if needed removing points.)

import java.io.FileWriter;
import java.io.IOException;
import java.lang.Math;
import java.util.ArrayList;
import java.util.Iterator;
import java.util.List;
public class PlaneRANSAC {
    PointCloud pointCloudPlaneRANSAC;
    double epsilonValue;
    Plane3D plane3DPlaneRANSAC;
    double support; 
    Point3D point1;
    Point3D point2;
    Point3D point3;
    public PlaneRANSAC(PointCloud pc){ //Constructor
        boolean flag = true;
        this.pointCloudPlaneRANSAC = pc; 
        this.support = 0; 
        Point3D point1 = pointCloudPlaneRANSAC.getPoint();
        Point3D point2 = pointCloudPlaneRANSAC.getPoint();
        Point3D point3 = pointCloudPlaneRANSAC.getPoint();
        while(flag){
            if(point1.equals(point3) || point2.equals(point3) || point1.equals(point2) ){ //Assumption: Points should be different or else plane would be 0
                point1 = pointCloudPlaneRANSAC.getPoint();
                point2 = pointCloudPlaneRANSAC.getPoint();
                point3 = pointCloudPlaneRANSAC.getPoint();
                flag = true;
            }
            else{
                flag = false;
            }
        }
        this.point1 = point1;
        this.point2 = point2;
        this.point3 = point3;
        Plane3D planeTestPlane3D = new Plane3D(point1,point2,point3);
        this.plane3DPlaneRANSAC = planeTestPlane3D;
    }
    public void setEps(double eps){ //Setter for epsilon
        this.epsilonValue = eps;
    }
    public double getEps(){ //Getter for epsilon
        return epsilonValue;
    }
    public int getNumberOfIterations(double confidence, double percentageOfPointsOnPlane){ //Get Approximate Number of Iterations needed to find dominant plane
        double temp = Math.ceil(Math.log(1-confidence)/Math.log(1-Math.pow(percentageOfPointsOnPlane, 3)));
        return (int) temp;
        }
    public void run(int numberOfIterations, String filename){ //Run RANSAC (ALL PRINT FUNCTIONS ARE FOR TRACING THE CODE | MORE THAN WELCOME TO UNCOMMENT AND SEE)
        int countPoints = 0;
        // int countSize = 0;
        List<Point3D> dominantPlane = new ArrayList<Point3D>();
        List<Point3D> currentDominantPlane = new ArrayList<Point3D>();
        List<Point3D> dominantPlaneAllPoints = new ArrayList<Point3D>();
        for(int threeMostDominantPlanes = 1; threeMostDominantPlanes < 4; threeMostDominantPlanes++){
            // System.out.println("RUN: " + threeMostDominantPlanes);
            // System.out.println("__________________________________________________________");
            for(int j = 0; j < numberOfIterations; j++){
                PlaneRANSAC tempPlane = new PlaneRANSAC(pointCloudPlaneRANSAC); //Creates New Plane | Adds points whenever distance is lower than epsilonValue and resets if new support is found
                Iterator<Point3D> iteratorTemp = pointCloudPlaneRANSAC.iterator();
                while(iteratorTemp.hasNext()){
                    Point3D tempPoint =  iteratorTemp.next();
                    double tempDistance = tempPlane.plane3DPlaneRANSAC.getDistance(tempPoint);
                    // countSize++;
                    if(tempDistance < epsilonValue){
                        countPoints++;
                        dominantPlane.add(tempPoint);
                    }
                }
                // System.out.println("Run: " + threeMostDominantPlanes + " Test Number: " + j);
                // System.out.println("Point 1: " + tempPlane.point1);
                // System.out.println("Point 2: " + tempPlane.point2);
                // System.out.println("Point 3: " + tempPlane.point3);
                // System.out.println("Plane: " + tempPlane.plane3DPlaneRANSAC.toString());
                // System.out.println("Points: " + countPoints);
                // System.out.println("Size: " + countSize);
                // System.out.println("==========================================================");
                // countSize = 0;
                if(countPoints > support){
                    currentDominantPlane.clear();
                    support = countPoints;
                    currentDominantPlane.addAll(dominantPlane); 
                    // System.out.println("Size: " + currentDominantPlane.size());
                    dominantPlane.clear();
                }
                dominantPlane.clear();
                countPoints = 0;
            }
            // System.out.println(currentDominantPlane);
            for(int withoutDominantPoints = 0; withoutDominantPoints < currentDominantPlane.size(); withoutDominantPoints++ ){ //Adding all dominant points to a list in all 3 runs
                if(!dominantPlaneAllPoints.contains(currentDominantPlane.get(withoutDominantPoints))){ //Only add to list if dominant Plane point isn't in the list
                    dominantPlaneAllPoints.add(currentDominantPlane.get(withoutDominantPoints));
                }
            }
            try {
                String replaceIndexForFile = String.valueOf(threeMostDominantPlanes); //Write p1,p2,p3 Files
                char charReplaceIndexForFile = replaceIndexForFile.charAt(0);
                String tempFileName = filename;
                tempFileName = tempFileName.replace('X', charReplaceIndexForFile);
                String stringFileName = tempFileName;
                FileWriter dominantPlanePointFile = new FileWriter(stringFileName);
                Iterator<Point3D> iteratorForpX = pointCloudPlaneRANSAC.iterator();
                dominantPlanePointFile.write("x y z");
                dominantPlanePointFile.write("\n");
                while(iteratorForpX.hasNext()){
                    Point3D tempPoint =  iteratorForpX.next();
                    if(currentDominantPlane.contains(tempPoint)){
                        dominantPlanePointFile.write(tempPoint.getX() + " " + tempPoint.getY() + " " + tempPoint.getZ());
                        dominantPlanePointFile.write("\n");
                    }
                }
                System.out.println("Successfully Wrote PointCloud into " + stringFileName);
                dominantPlanePointFile.close();
            } catch (IOException e) {
                System.out.println("An error occurred.");
                e.printStackTrace();
              }

            dominantPlane.clear();
            // System.out.println("Current Support: " + support);
            // System.out.println("==========================================================");
            support = 0;
        }  
        
        for(int p0 = 0; p0 < dominantPlaneAllPoints.size(); p0++){ //Remove all points that are dominant out of 3 runs
            if(pointCloudPlaneRANSAC.pointcloud.contains(dominantPlaneAllPoints.get(p0))){
                pointCloudPlaneRANSAC.pointcloud.remove(dominantPlaneAllPoints.get(p0));
            }
        }

        String tempFileName = filename; //Writes p0 file with save() function from PointCloud.java
        tempFileName = tempFileName.replace('X', '0');
        String stringFileName = tempFileName;
        pointCloudPlaneRANSAC.save(stringFileName);
}
    public static void main(String[] args){ //Test RANSAC function
        PointCloud pc1 = new PointCloud("PointCloud1.xyz"); 
        PlaneRANSAC testPlane1 = new PlaneRANSAC(pc1);
        testPlane1.setEps(0.1);
        testPlane1.run(testPlane1.getNumberOfIterations(0.99, 0.05),"PointCloud1_pX.xyz");

        PointCloud pc2 = new PointCloud("PointCloud2.xyz"); 
        PlaneRANSAC testPlane2 = new PlaneRANSAC(pc2);
        testPlane2.setEps(0.1);
        testPlane2.run(testPlane2.getNumberOfIterations(0.99, 0.05),"PointCloud2_pX.xyz");

        PointCloud pc3 = new PointCloud("PointCloud3.xyz"); 
        PlaneRANSAC testPlane3 = new PlaneRANSAC(pc3);
        testPlane3.setEps(0.1);
        testPlane3.run(testPlane3.getNumberOfIterations(0.99, 0.05),"PointCloud3_pX.xyz");

    }
}
