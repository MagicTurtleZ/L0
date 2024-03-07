document.getElementById("submit-btn").addEventListener("click", function() {
    var orderId = document.getElementById("order-id").value;

    setTimeout(function() {
        fetch('http://localhost:8080/order', {
            method: 'POST',
            body: JSON.stringify({ orderId: orderId }),
            headers: {
                'Content-Type': 'application/json'
            }
        })
        .then(response => response.json())
        .then(data => {
            var resultMessage = `Order with ID ${orderId}:`;
            document.getElementById("result").innerText = resultMessage;
            
            // Вывод полученных данных в поле order-info
            var orderInfo = `${data.orderInfo}`;
            var status = `${data.status}`;
            if (status == "ERROR") {
                document.getElementById("order-info").innerText = "The record was not found";
            } else {
                document.getElementById("order-info").innerText = orderInfo;
            }
        })
        .catch(error => {
            console.error('Error:', error);
        });
    });
});