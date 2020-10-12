package gospline

//
// Gospline package for calculating spline-based regression models in Go
// AUTHOR: TRENT HENDERSON, 12 OCTOBER 2020
//

//------------------------------------------------------------------

// Import packages and dependencies needed

import (
  "fmt"
  "math"
  "gonum.org/v1/gonum/floats"
  "gonum.org/v1/plot"
  "gonum.org/v1/plot/plotter"
  "gonum.org/v1/plot/plotutil"
  "gonum.org/v1/plot/vg"
)

//------------------------------------------------------------------
// Instantiate data to use in the model
// x = input variable
// y = response variable
// NOTE: MODEL CURRENTLY ONLY SUPPORTS BIVARIATE RELATIONSHIPS
// IT IS INTENDED THAT A THEORETICALLY INFINITE NUMBER OF INPUT VARIABLES
// WILL BE ADDED

type Init struct {
  x, y []float64
}

//------------------------------------------------------------------

// Define polynomial spline regression function using basis functions
//
// PARAMETERS:
// x = predictor variable to use in the model
// y = response variable to use in the model
// k = number of knots to fit - controls number of basis functions
// l = order of polynomial (suggested option is l = 3 for a "cubic")

func gam(*Init, k int, l int) [] float64 {

  // Conditions to cancel loop

  if len(x) < 5 {
    fmt.Println("Not enough data to compute meaningful basis functions.")
    } else if len(y) < 5 {
      fmt.Println("Not enough data to compute meaningful basis functions.")
      } else if (float64(len(x)) / float64(k)) < float64(3) {
        fmt.Println("Not enough degrees of freedom. Please specify less knots.")
        } else if l < 1 {
          fmt.Println("Please specify an integer polynomial degree between 1 and 5 for best results.")
          } else if l > 5 {
            fmt.Println("Please specify an integer polynomial degree between 1 and 5 for best results.")
            } else if len(x) != len(y) {
              fmt.Println("X and Y variables should be the same length.")
              } else {

    // Define number of knots

    the_min := floats.Min(x)
    the_max := floats.Max(x)

    step := (the_max - the_min) / (float64(k) - 1)
    var knots []float64

    for f := 1; f < k; f++ {
      knots[f] = the_min + step * float64(f)
    }

    // Calculate basis values

    var bs []float64

    for g := 0; g < len(knots); g++ {
      for i := 0; i < len(x); i++ {
        if x >= float64(g) {
          bs[i] = math.Pow((x[i] - knots[g]), float64(l))
          } else {
            bs[i] = 0
        }
      }
    }

    // Get linear regression coefficients

     x_mean := mean(bs)
     y_mean := mean(y)

     var numerator []float64
     var denominator []float64

     for i := 0; i < len(bs); i++ {
       numerator[i] += (bs[i]- x_mean) * (y[i] - y_mean)
       denominator[i] += (bs[i]- x_mean) * (bs[i]- x_mean)
     }

     var b1 []float64

     for i := 0; i < len(numerator); i++ {
       b1[i] += numerator[i] / denominator[i]
     }

    // Multiply basis functions by regression coefficients

    var results []float64

    for i := 0; i < len(b1); i++ {
      results[i] += bs[i] * b1[i]
    }
    return results
  }
}

//------------------------------------------------------------------
// Function to compute a mean
//
// PARAMETERS:
// arr = slice from which to calculate the mean of

func mean(arr [] float64) float64 {
   sum := 0.0
   for _, v := range arr {
      sum += v
   }
   return sum / float64(len(arr))
}
