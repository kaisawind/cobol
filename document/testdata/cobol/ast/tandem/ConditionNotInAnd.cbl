 IDENTIFICATION DIVISION.
 PROGRAM-ID. CONDTN.
 PROCEDURE DIVISION.   
    IF SOME-DATA IN SOME-OTHER-DATA NOT = ('A' AND 'B')
       STOP RUN
    END-IF.