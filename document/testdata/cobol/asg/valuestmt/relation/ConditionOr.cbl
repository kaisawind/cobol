 IDENTIFICATION DIVISION.
 PROGRAM-ID. COND.
 DATA DIVISION.
 WORKING-STORAGE SECTION.
 77 SOMECOND1 PICTURE IS 9(1).
 77 SOMECOND2 PICTURE IS 9(1).
 PROCEDURE DIVISION.
    IF SOMECOND1 OR NOT SOMECOND2 THEN END-IF.