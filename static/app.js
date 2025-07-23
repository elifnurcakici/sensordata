const ws = new WebSocket("ws://localhost:8080/ws");

        const tempCtx = document.getElementById('tempChart').getContext('2d');
        const humidityCtx = document.getElementById('humidityChart').getContext('2d');
        const pressureCtx = document.getElementById('pressureChart').getContext('2d');

        const gradients = {
            temp: createGradient(tempCtx, 'rgba(255,99,132,0.4)', 'rgba(255,99,132,0.8)'),
            humidity: createGradient(humidityCtx, 'rgba(54,162,235,0.4)', 'rgba(54,162,235,0.8)'),
            pressure: createGradient(pressureCtx, 'rgba(75,192,192,0.4)', 'rgba(75,192,192,0.8)')
        };

        const tempChart = createChart(tempCtx, 'Sıcaklık (°C)', gradients.temp, 'rgb(255,99,132)');
        const humidityChart = createChart(humidityCtx, 'Nem (%)', gradients.humidity, 'rgb(54,162,235)');
        const pressureChart = createChart(pressureCtx, 'Basınç (hPa)', gradients.pressure, 'rgb(75,192,192)');

        ws.onmessage = function(event) {
            const msg = JSON.parse(event.data);
            const timeLabel = new Date(msg.data.created_at).toLocaleTimeString();

            if (msg.type === "temperature") {
                updateChart(tempChart, timeLabel, msg.data.value);
                updateTable("temperatureTable", msg.data);
            }
            if (msg.type === "humidity") {
                updateChart(humidityChart, timeLabel, msg.data.value);
                updateTable("humidityTable", msg.data);
            }
            if (msg.type === "pressure") {
                updateChart(pressureChart, timeLabel, msg.data.value);
                updateTable("pressureTable", msg.data);
            }
        };

        function createGradient(ctx, color1, color2) {
            const gradient = ctx.createLinearGradient(0, 0, 0, 400);
            gradient.addColorStop(0, color1);
            gradient.addColorStop(1, color2);
            return gradient;
        }

        function createChart(ctx, label, gradient, borderColor) {
            return new Chart(ctx, {
                type: 'line',
                data: { labels: [], datasets: [{ label, data: [], borderColor, backgroundColor: gradient, fill: true, tension: 0.4 }] },
                options: {
                    responsive: true,
                    plugins: { legend: { labels: { color: '#fff' } } },
                    scales: { x: { ticks: { color: '#fff' } }, y: { ticks: { color: '#fff' } } }
                }
            });
        }

        function updateChart(chart, label, value) {
            if (chart.data.labels.length > 20) {
                chart.data.labels.shift();
                chart.data.datasets[0].data.shift();
            }
            chart.data.labels.push(label);
            chart.data.datasets[0].data.push(value);
            chart.update();
        }

        function updateTable(tableId, data) {
            const table = document.getElementById(tableId).querySelector("tbody");
            const row = document.createElement("tr");
            row.innerHTML = `<td>${data.id}</td><td>${data.value.toFixed(2)}</td><td>${new Date(data.created_at).toLocaleTimeString()}</td>`;
            table.prepend(row);

            if (table.rows.length > 10) {
                table.deleteRow(10);
            }
        }

        window.onload = async function() {
            await loadInitialData("temperature");
            await loadInitialData("humidity");
            await loadInitialData("pressure");
        };

        async function loadInitialData(sensor) {
            const res = await fetch(`/data/${sensor}/all`);
            const data = await res.json();
            const table = document.getElementById(`${sensor}Table`).querySelector("tbody");
            table.innerHTML = "";

            data.slice(-10).reverse().forEach(d => {
                const row = document.createElement("tr");
                row.innerHTML = `<td>${d.id}</td><td>${d.value.toFixed(2)}</td><td>${new Date(d.created_at).toLocaleTimeString()}</td>`;
                table.appendChild(row);
            });

            const chart = sensor === "temperature" ? tempChart : (sensor === "humidity" ? humidityChart : pressureChart);
            chart.data.labels = data.map(d => new Date(d.created_at).toLocaleTimeString());
            chart.data.datasets[0].data = data.map(d => d.value);
            chart.update();
        }