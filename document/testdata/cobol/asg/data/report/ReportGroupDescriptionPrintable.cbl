 IDENTIFICATION DIVISION.
 PROGRAM-ID. REPGRPDESCVERT.
 DATA DIVISION.
    REPORT SECTION.
    RD REPORT1
       IS GLOBAL.
       01 SOMEDATANAME
          PICTURE IS 9(10)
          SIGN IS TRAILING SEPARATE
          JUST RIGHT
          BLANK WHEN ZERO
          LINE NUMBER IS 2
          COLUMN NUMBER IS 42
          SUM SOMEID , SOMEID2 , SOMEID3
             UPON SOMEID4 , SOMEID5
          USAGE IS DISPLAY-1
       .