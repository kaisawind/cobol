 IDENTIFICATION DIVISION.
 PROGRAM-ID. RECEIVEINTOSTMT.
 PROCEDURE DIVISION.
    RECEIVE SOMECD1 FROM THREAD SOMETHREAD1
       BEFORE TIME SOMETIME1
       WITH NO WAIT
       THREAD IN SOMETHREAD1
       SIZE IN SOMEID1
       STATUS IN SOMEID2.