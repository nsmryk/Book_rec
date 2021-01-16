var ctx = document.getElementById('myChart').getContext('2d');
var chart = new Chart(ctx, {
    // The type of chart we want to create
    type: 'line',

    // The data for our dataset
    data: {
        datasets: [{
            label: 'My First dataset',
            backgroundColor: 'rgb(255, 99, 132)',
            borderColor: 'rgb(255, 99, 132)',
            fill: false,
            data: [{
                x: '1995-12-17T00:00:00',
                y: 1
            }, {
                x: '1995-12-18T00:00:00',
                y: 10
            }, {
                x: '1995-12-20T00:00:00',
                y: 20
            }, {
                x: '1995-12-21T12:00:00',
                y: 30
            }, {
                x: '1996-01-01T00:00:00',
                y: 40
            }]
        }]
    },
    // Configuration options go here
    options: {
        scales: {
            xAxes: [{
                type: 'time',
                time: {
                    unit: 'day'
                }
            }]
        }
    }
});