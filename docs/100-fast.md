---
title: How much faster?
---

Let's see how much faster antibody is over antigen:

<div id="chart1"></div>
<div id="chart2"></div>


<script src="https://unpkg.com/frappe-charts@0.0.8/dist/frappe-charts.min.iife.js"></script>

<script type="text/javascript">
  let data1 = {
    labels: ["warmup"],
    datasets: [
      {
        title: "Antigen",
        values:  [72.85],
      },
      {
        title: "Antibody",
        values: [3.99]
      }
    ]
  };

  let chart1 = new Chart({
    parent: "#chart1", // or a DOM element
    title: "First load of 20 libraries (warmup):",
    data: data1,
    type: 'bar', // or 'line', 'scatter', 'pie', 'percentage'
    show_dots: 0,
    height: 250,
    x_axis_mode: 'tick',
    is_series: 1,
    format_tooltip_x: d => ('' + d).toUpperCase(),
    format_tooltip_y: d => d + 's'
  });
  let data2 = {
    labels: ["1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31", "32", "33", "34", "35", "36", "37", "38", "39", "40", "41", "42", "43", "44", "45", "46", "47", "48", "49", "50"],
    datasets: [
      {
        title: "Antigen",
        values: [2.57,2.71,2.11,2.27,2.47,2.34,2.27,2.47,2.36,2.27,2.42,2.24,2.53,2.85,2.22,2.49,2.41,2.32,2.37,2.20,2.16,2.46,2.42,2.41,2.34,2.26,2.31,2.25,2.17,2.51,2.49,2.11,2.26,2.32,2.31,2.27,2.11,2.21,2.59,2.26,2.27,2.30,2.26,2.86,2.44,2.41,2.20,2.15,2.16,2.11]
      },
      {
        title: "Antibody",
        values: [0.08,0.08,0.08,0.08,0.07,0.08,0.08,0.08,0.08,0.08,0.08,0.08,0.09,0.08,0.08,0.08,0.08,0.08,0.08,0.08,0.08,0.08,0.08,0.08,0.08,0.08,0.08,0.07,0.08,0.08,0.08,0.08,0.08,0.08,0.07,0.08,0.08,0.08,0.08,0.08,0.08,0.08,0.08,0.08,0.08,0.08,0.08,0.08,0.08,0.08]
      }
    ]
  };

  let chart2 = new Chart({
    parent: "#chart2", // or a DOM element
    title: "Times taken to load the same 20 libraries 50 times (after the warmup)",
    data: data2,
    type: 'line', // or 'line', 'scatter', 'pie', 'percentage'
    show_dots: 0,
    height: 250,
    x_axis_mode: 'tick',
    is_series: 1,
    format_tooltip_x: d => ('run ' + d).toUpperCase(),
    format_tooltip_y: d => d + 's'
  });
</script>
