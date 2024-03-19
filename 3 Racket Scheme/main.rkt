#lang racket
(require racket/file)

(define (readXYZ fileIn)
  (let ((sL (map (lambda s (string-split (car s)))
                 (cdr (file->lines fileIn)))))
    (map (lambda (L)
           (map (lambda (s)
                  (if (eqv? (string->number s) #f)
                      s
                      (string->number s))) L)) sL)))


;(readXYZ "Point_Cloud_1_No_Road_Reduced.xyz")
(define (plane P1 P2 P3)
  (let ((p1x (car P1))
        (p1y (car (cdr P1)))
        (p1z (car (cdr (cdr P1))))
        (p2x (car P2))
        (p2y (car (cdr P2)))
        (p2z (car (cdr (cdr P2))))
        (p3x (car P3))
        (p3y (car (cdr P3)))
        (p3z (car (cdr (cdr P3)))))
    (let ((ax (- p2x p1x))
          (ay (- p2y p1y))
          (az (- p2z p1z))
          (bx (- p3x p1x))
          (by (- p3y p1y))
          (bz (- p3z p1z))
          (px p1x)
          (py p1y)
          (pz p1z))
      (let ((crossProduct (list (- (* ay bz) (* az by))
                                 (- (* az bx) (* ax bz))
                                 (- (* ax by) (* ay bx))
                                 (-
                                 (* -1 (* (- (* ay bz) (* az by)) px))
                                 (* (- (* az bx) (* ax bz)) py)
                                 (* (- (* ax by) (* ay bx)) pz))
                                 )))
        ;(list p1x p1y p1z p2x p2y p2z p3x p3y p3z
              ;(car crossProduct) 
              ;(cadr crossProduct) 
              ;(caddr crossProduct) 
              ;(cadddr crossProduct))
        (list (car crossProduct) 
              (cadr crossProduct) 
              (caddr crossProduct) 
              (cadddr crossProduct))        
        ))))

(define (count-up count)
  (lambda ()
    (let ((new-count (+ count 1)))
      (set! count new-count)
      new-count)))

(define (loop lst plane eps)
  (let ((support-points '())
        (count (count-up 0)))
    (define (helper lst)
      (cond ((null? lst) support-points)
            (else
             (let ((currentPX (car (car lst))) 
                   (currentPY (cadr (car lst))) 
                   (currentPZ (caddr (car lst))))
               (let ((distance (if (zero? (cadddr plane))
                                    #f
                                    (/ (+ (* (abs (car plane)) currentPX)
                                          (* (abs (cadr plane)) currentPY)
                                          (* (abs (caddr plane)) currentPZ))
                                       (cadddr plane)))))
                 (if distance
                     (if (< distance eps)
                         (set! support-points (cons (car lst) support-points))
                         #f)
                     #f))
               (helper (cdr lst))))))
    (helper lst)
    support-points))

(define (support plane points eps)
  (let ((support-points (loop points plane eps)))
    (list (length support-points) plane)))

(define (dominantPlane Ps k eps)
  (let loop ((i 0) (best-support 0) (best-plane '()) (best-supports '()) (best-planes '()))
    (cond ((>= i k) (list (car best-supports) (car(car best-planes))  (cadr(car best-planes)) (caddr(car best-planes)) (cadddr(car best-planes))))
          (else
           (let ((P1 (list-ref Ps (random (length Ps))))
                 (P2 (list-ref Ps (random (length Ps))))
                 (P3 (list-ref Ps (random (length Ps)))))
             (let ((currentPlane (plane P1 P2 P3))
                   (currentSupport (car(support (plane P1 P2 P3) Ps eps))))
               (if (> currentSupport best-support)
                   (loop (+ i 1) currentSupport currentPlane (cons currentSupport best-supports) (cons currentPlane best-planes))
                   (loop (+ i 1) best-support best-plane best-supports best-planes))
               ))))
    )
  )

(define (ransacNumberOfIteration confidence percentage)
  (ceiling (/ (ceiling (log (- 1 confidence)))  (log (- 1 (expt percentage 3)))))
  )

(define (planeRANSAC filename confidence percentage eps)
  (let ((Ps (readXYZ filename)))
    (dominantPlane Ps (ransacNumberOfIteration confidence percentage) eps)
          ))

;confidence percentage eps
(planeRANSAC "Point_Cloud_1_No_Road_Reduced.xyz" 0.99 0.05 0.2)
(planeRANSAC "Point_Cloud_2_No_Road_Reduced.xyz" 0.99 0.05 0.2)
(planeRANSAC "Point_Cloud_3_No_Road_Reduced.xyz" 0.99 0.05 0.2)