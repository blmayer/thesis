sdl[vals_List] := Block[{counts, Hmax, H, delta}, 
     counts = (Counts[#1]/Length[vals] & )[vals]; Hmax = Log[Length[counts]]; 
      H = N[-Total[(#1*Log[#1] & ) /@ counts]]; delta = H/Hmax; 
      N[{4*delta*(1 - delta), delta}]]
