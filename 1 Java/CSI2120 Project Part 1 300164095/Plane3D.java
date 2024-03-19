/**
 * Student Name: Axel Tang
 * Student ID: 300164095
 * CSI 2120 Project: Part 1 
 */

import java.lang.Math;
public class Plane3D {
    double plane3DX, plane3DY,plane3DZ,plane3DC;

    public Plane3D(Point3D p1, Point3D p2, Point3D p3){ //Constructor
        double ax = p2.getX() - p1.getX();
        double ay = p2.getY() - p1.getY();
        double az = p2.getZ() - p1.getZ();

        double bx = p3.getX() - p1.getX();
        double by = p3.getY() - p1.getY();
        double bz = p3.getZ() - p1.getZ();

        double px = p1.getX();
        double py = p1.getY();
        double pz = p1.getZ();

        double[] crossProduct = new double[4];
        crossProduct[0] = (ay)*(bz) - (az)*(by);
        crossProduct[1] = (az)*(bx) - (ax)*(bz);
        crossProduct[2] = (ax)*(by) - (ay)*(bx);
        crossProduct[3] = (-crossProduct[0] * px) - (crossProduct[1] * py) - (crossProduct[2] * pz);

        this.plane3DX = crossProduct[0];
        this.plane3DY = crossProduct[1];
        this.plane3DZ = crossProduct[2];
        this.plane3DC = crossProduct[3];
    }
    public Plane3D(double a, double b, double c, double d){ //Constructor
        this.plane3DX = a;
        this.plane3DY = b;
        this.plane3DZ = c;
        this.plane3DC = d;
        // System.out.println(a + " " + b + " " + c + " " + d);
    }
    public double getDistance(Point3D pt){ //Get Distance Function
        double distance = (Math.abs(plane3DX*pt.getX() + plane3DY*pt.getY() + plane3DZ*pt.getZ() + plane3DC))/(Math.sqrt(Math.pow(plane3DX,2) + Math.pow(plane3DY,2) + Math.pow(plane3DZ,2)));
        return distance;
    }
    // public static void main(String[] args){ //Test Function
    //     Point3D p1 = new Point3D(1,2,3);
    //     Point3D p2 = new Point3D(2,10,12);
    //     Point3D p3 = new Point3D(7,8,9);
    //     Plane3D test = new Plane3D(p1, p2, p3);
    //     Plane3D testPLane = new Plane3D(1, 2, 3, 4); //Test Other Constructor
    //     System.out.println("Distance: " + test.getDistance(new Point3D(4, 10, 9)));
    // }
    // @Override
    // public String toString() {
    //     return this.plane3DX + " + " + this.plane3DY + " + " +  this.plane3DZ + " + " + this.plane3DC;
    // }
}
