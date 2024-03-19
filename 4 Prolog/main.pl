%Given Code to read the file in prolog 
read_xyz_file(File, Points) :-
    open(File, read, Stream),
    read_xyz_points(Stream,Points),
    close(Stream).

read_xyz_points(Stream, []) :-
    at_end_of_stream(Stream).

read_xyz_points(Stream, [Point|Points]) :-
    \+ at_end_of_stream(Stream),
read_line_to_string(Stream,L), split_string(L, "\t", "\s\t\n",XYZ), convert_to_float(XYZ,Point),
    read_xyz_points(Stream, Points).

convert_to_float([],[]).

convert_to_float([H|T],[HH|TT]) :-
    atom_number(H, HH),
    convert_to_float(T,TT).

%Generates 3 random points from the list
random3points(Points, Point3) :-
    length(Points, Length),
    random(0, Length, Index1),
    random(0, Length, Index2),
    random(0, Length, Index3),
    nth0(Index1, Points, Element1),
    nth0(Index2, Points, Element2),
    nth0(Index3, Points, Element3),
    Point3 = [Element1,Element2,Element3].

%Generates a plane equation of ax + by + cz = d based off the 3 points given
plane(Point3, Plane) :-
    nth0(0,Point3,P1),
    nth0(1,Point3,P2),
    nth0(2,Point3,P3),
    nth0(0,P1,P1X),
    nth0(1,P1,P1Y),
    nth0(2,P1,P1Z),
    nth0(0,P2,P2X),
    nth0(1,P2,P2Y),
    nth0(2,P2,P2Z),
    nth0(0,P3,P3X),
    nth0(1,P3,P3Y),
    nth0(2,P3,P3Z),
    AX is P2X - P1X,
    AY is P2Y - P1Y,
    AZ is P2Z - P1Z,
    BX is P3X - P1X,
    BY is P3Y - P1Y,
    BZ is P3Z - P1Z,
    CrossProduct0 is AY*BZ - AZ*BY,
    CrossProduct1 is AZ*BX - AX*BZ,
    CrossProduct2 is AX*BY - AZ*BX,
    CrossProduct3 is (-CrossProduct0 * P1X) - (CrossProduct1 * P1Y) - (CrossProduct2 * P1Z),
    Plane = [CrossProduct0,CrossProduct1,CrossProduct2,CrossProduct3].

%Initialize loopList Function.
loopList([], _, Counter, _, _, _, _, Counter).

%loopList is a function which iteratively loops over the entire PointCloud given and is used to get the 
%Support Count which compares every point's towrads plane's distance
% with eps value 
loopList([Head|Tail], Eps, Counter, Plane3DX, Plane3DY, Plane3DZ, Plane3DC, Result) :-
        nth0(0, Head, PTX),
        nth0(1, Head, PTY),
        nth0(2, Head, PTZ),
        Distance is abs(Plane3DX*PTX + Plane3DY*PTY + Plane3DZ*PTZ +Plane3DC)/sqrt(Plane3DX**2 + Plane3DY**2 + Plane3DZ**2),
        (Distance < Eps ->
            Counter1 is Counter + 1
        ;   
            Counter1 = Counter
        ),
        loopList(Tail, Eps, Counter1, Plane3DX, Plane3DY, Plane3DZ, Plane3DC, Result).

%Returns true or false depending if N is the correct value of supporting points. 
support(Plane, Points, Eps, N) :-
    nth0(0, Plane, Plane3DX),
    nth0(1, Plane, Plane3DY),
    nth0(2, Plane, Plane3DZ),
    nth0(3, Plane, Plane3DC),
    loopList(Points, Eps, 0, Plane3DX, Plane3DY, Plane3DZ, Plane3DC, Temp),
    Test = Temp,
    (Test =:= N ->
        write('SupportCount = true ')
    ;   
        write('SupportCount = false ')
    ).
    
%Returns true or false depending if N is the correct value of number of iterations needed. 
ransac-number-of-iterations(Confidence, Percentage, N) :-
    Test = ceil(log(1 - Confidence) / log(1 - (Percentage**3))),
    (Test =:= N ->
        write('Num Of Iterations = true ')
    ;   
        write('Num Of Iterations = false ')
    ).

% ?- read_xyz_file("Point_Cloud_1_No_Road_Reduced.xyz", X), test3Point(X, N).
% Will show N = 1 

test3Point(Point,1) :- 
    random3points(Point , Point3),
    nth0(0,Point3,P1),
    nth0(0,Point3,P2),
    nth0(0,Point3,P3),
    member(P1, Point),
    member(P2, Point),
    member(P3, Point).
% ?- read_xyz_file("Point_Cloud_1_No_Road_Reduced.xyz", X), test3Point(X, 1).
% Will Give The Output of X 

% ?- read_xyz_file("Point_Cloud_1_No_Road_Reduced.xyz", X), test3Point(X, 2).
% Will Give An Error ( Which is false )

test3Point(Point,2) :- 
    random3points(Point , Point3),
    nth0(0,Point,P1),
    nth0(1,Point,P2),
    nth0(2,Point,P3),
    member(P2, Point3),
    member(P1, Point3),
    member(P3, Point3).

%Support Testing:
% Plane = [0.9907113715583863, 68.75622472210522, -65.75581315283364, -779.007408767287]
% Eps = 0.02
% SupportCount will be 16.

%Test Case 1 will be true , Test Case 2 will be false.
%testSupport(N) , gives N = 1.
%testSupport(1), returns SupportCount = true, testSuport(2), returns SupportCount = false. | Both returns value true for indication of test case sucessfully running.

testSupport(1) :- 
    Plane = [0.9907113715583863, 68.75622472210522, -65.75581315283364, -779.007408767287],
    Eps = 0.02,
    read_xyz_file("Point_Cloud_1_No_Road_Reduced.xyz",X),
    support(Plane,X,Eps,16).

testSupport(2) :- 
    Plane = [0.9907113715583863, 68.75622472210522, -65.75581315283364, -779.007408767287],
    Eps = 0.02,
    read_xyz_file("Point_Cloud_1_No_Road_Reduced.xyz",X),
    support(Plane,X,Eps,15).

%Ransac Iterations Testing:
% Confidence = 0.99
% Percentage = 0.05 
% Number Of Iteartions will be 36840.

%Test Case 1 will be true , Test Case 2 will be false.
%testSupport(N) , gives N = 1.
%testSupport(1), returns Num Of Iterations = true, testSuport(2), returns Num Of Iterations = false. | Both returns value true for indication of test case sucessfully running.
testRansacIteration(1) :-
    Confidence = 0.99,
    Percentage = 0.05 ,
    ransac-number-of-iterations(Confidence, Percentage, 36840).
testRansacIteration(2) :-
    Confidence = 0.99,
    Percentage = 0.05 ,
    ransac-number-of-iterations(Confidence, Percentage, 36842).

%Test Function to read all files.
% Run with 
%?- testReadAllXYZFiles(PC1,PC2,PC3).
testReadAllXYZFiles(PC1,PC2,PC3) :- 
    read_xyz_file("Point_Cloud_1_No_Road_Reduced.xyz",PC1),
    read_xyz_file("Point_Cloud_2_No_Road_Reduced.xyz",PC2),
    read_xyz_file("Point_Cloud_3_No_Road_Reduced.xyz",PC3).
