(* Created with the Wolfram Language : www.wolfram.com *)
{visibleQ, horizontalQ, naturalVisibility, horizontalVisibility, plotModel, 
 plotLinearModel}
visibleQ[a_Integer, b_Integer, ts_List] := 
    Catch[Do[If[ !ts[[x,2]] < ts[[b,2]] + (ts[[a,2]] - ts[[b,2]])*
            ((ts[[b,1]] - ts[[x,1]])/(ts[[b,1]] - ts[[a,1]])), Throw[False]], 
       {x, a + 1, b - 1}]; Throw[True]]
 
ts = {{1, 1}, {7, 2}, {10, 4}, {13, 2}, {19, 2}, {23, 2}, {28, 6}, {31, 2}, 
     {32, 6}, {44, 6}, {49, 3}, {68, 6}, {70, 8}, {79, 2}, {82, 4}, {86, 4}, 
     {91, 4}, {94, 4}}
 
horizontalQ[a_Integer, b_Integer, ts_List] := 
    Catch[Do[If[ !(ts[[x,2]] < ts[[a,2]] && ts[[x,2]] < ts[[b,2]]), 
        Throw[False]], {x, a + 1, b - 1}]; Throw[True]]
 
naturalVisibility[series_] := Block[{visibles}, 
     visibles = {}; SetSharedVariable[visibles]; 
      ParallelDo[Do[If[visibleQ[i, j, series], AppendTo[visibles, 
          series[[i,1]] -> series[[j,1]]]], {j, i + 1, Length[series]}], 
       {i, 1, Length[series] - 1}]; Return[visibles]]
 
horizontalVisibility[series_] := Block[{visibles}, 
     visibles = {}; SetSharedVariable[visibles]; 
      ParallelDo[Do[If[horizontalQ[i, j, series], AppendTo[visibles, 
          series[[i,1]] -> series[[j,1]]]], {j, i + 1, Length[series]}], 
       {i, 1, Length[series] - 1}]; Return[visibles]]
 
plotModel[count_, len_] := Block[{countrel, nlm, x, C, \[Gamma]}, 
     countrel = Table[{x, Lookup[count, x]/len}, {x, Keys[count]}]; 
      nlm = NonlinearModelFit[countrel, C/x^\[Gamma], {C, \[Gamma]}, x]; 
      {Apply[{#1, Differences[#2][[1]]/2} & , 
        nlm["ParameterConfidenceIntervalTableEntries"][[All,{1, 3}]], {1}], 
       Show[ListPlot[countrel, PlotRange -> All, AxesOrigin -> {0, 0}], 
        Plot[nlm[x], {x, 0, Max[Keys[count]]}, PlotRange -> All, 
         AxesOrigin -> {0, 0}], AxesLabel -> {Style["k", 38, Italic], 
          Style["P(k)", 38, Italic]}, LabelStyle -> Large, 
        ImageSize -> Large]}]
 
plotLinearModel[count_, len_] := Block[{countrel, nlm, x, points}, 
     countrel = Table[{x, Lookup[count, x]/len}, {x, Keys[count]}]; 
      points = Log[countrel]; lm = LinearModelFit[points, x, x]; 
      Show[ListPlot[points, PlotRange -> {{0, 8.1}, All}, 
        AxesOrigin -> {0, 0}], Plot[lm[x], {x, 0, 8.1}, 
        PlotRange -> {{0, 8.1}, All}, AxesOrigin -> {0, 0}], 
       AxesLabel -> {Style["k", 38, Italic], Style["P(k)", 38, Italic]}, 
       LabelStyle -> Large, ImageSize -> Large, AspectRatio -> 0.8]]
