/**
 * Student Name: Axel Tang
 * Student ID: 300164095
 * CSI 2120 Project: Part 1 
 */
public class Point3D{
    double x,y,z;
    public Point3D(double x,double y,double z){ //Constructor
        this.x = x;
        this.y = y; 
        this.z = z;
    }
    public double getX(){ //Get Methods
        return x;
    }
    public double getY(){
        return y;
    }
    public double getZ(){
        return z;
    }
    // @Override
    // public String toString() {
    //     return "X: " + this.x + " | " + 
    //     "Y: " + this.y + " | " +
    //     "Z: " +  this.z;
    // }
}