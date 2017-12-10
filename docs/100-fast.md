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
        values:  [44.40],
      },
      {
        title: "Antibody",
        values: [10.51]
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
        values: [0.37,0.34,0.34,0.35,0.36,0.36,0.44,0.42,0.44,0.35,0.34,0.33,0.33,0.34,0.33,0.33,0.33,0.36,0.39,0.39,0.36,0.35,0.38,0.33,0.34,0.34,0.35,0.48,0.39,0.39,0.35,0.36,0.36,0.41,0.43,0.46,0.35,0.35,0.34,0.35,0.36,0.37,0.34,0.33,0.33,0.33,0.33,0.35,0.32,0.34]
      },
      {
        title: "Antibody",
        values: [0.15,0.14,0.15,0.15,0.15,0.15,0.16,0.18,0.15,0.16,0.16,0.26,0.22,0.16,0.18,0.17,0.16,0.17,0.15,0.14,0.16,0.15,0.15,0.18,0.18,0.14,0.14,0.28,0.14,0.14,0.14,0.14,0.15,0.15,0.14,0.15,0.16,0.18,0.15,0.15,0.17,0.14,0.15,0.15,0.15,0.14,0.14,0.15,0.15,0.16]
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
  chart2.show_averages();
</script>

Data from [getantibody/speed](https://github.com/getantibody/speed) repository.
