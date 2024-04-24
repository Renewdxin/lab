document.getElementById('loginForm').addEventListener('submit', function(event) {
    event.preventDefault();
    const username = this.username.value;
    const password = this.password.value;

    fetch('http://127.0.0.1:8080/auth/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            id: username,
            password: password
        })
    })
        .then(response => response.json())
        .then(data => {
            console.log('Success:', data);
            if (data.success) {
                alert('Login successful!');
                // 这里可以进行页面跳转或其他操作
            } else {
                alert('Login failed: ' + data.message);
            }
        })
        .catch((error) => {
            console.error('Error:', error);
            alert('Login failed: Network error or server is down');
        });
});

document.getElementById('registerForm').addEventListener('submit', function(event) {
    event.preventDefault();
    const username = this.username.value;
    const email = this.email.value;
    const password = this.password.value;

    fetch('http://127.0.0.1:8080/auth/register', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            username: username,
            email: email,
            password: password
        })
    })
        .then(response => response.json())
        .then(data => {
            console.log('Success:', data);
            if (data.success) {
                alert('Registration successful!');
                // 这里可以进行页面跳转或其他操作
            } else {
                alert('Registration failed: ' + data.message);
            }
        })
        .catch((error) => {
            console.error('Error:', error);
            alert('Registration failed: Network error or server is down');
        });
});
