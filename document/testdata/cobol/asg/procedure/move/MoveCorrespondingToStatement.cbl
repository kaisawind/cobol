 IDENTIFICATION DIVISION.
 PROGRAM-ID. MOVECORRTOSTMT.
 DATA DIVISION.
 WORKING-STORAGE SECTION.
 77 SOME-TEXT PICTURE IS X(9).
 77 SOME-TEXT2 PICTURE IS X(9).
 PROCEDURE DIVISION.
     MOVE CORRESPONDING SOME-TEXT TO SOME-TEXT2.
     MOVE CORR SOME-DATA (SOME-INDEX) TO SOME-DATA2.