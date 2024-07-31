<template>
    <div>
        <div class="container text-center mt-5">
            <h1>{{ isLogin ? 'Login' : 'Register' }}</h1>
            <form @submit.prevent="handleSubmit">
                <div class="mb-3">
                <label for="phone"  class="form-label widt">Phone</label>
                <input type="text" v-model="phone" class="form-control" id="phone" required>
                </div>
                <div class="mb-3">
                <label for="password"  class="form-label">Password</label>
                <input type="password" v-model="password" class="form-control" id="password" required>
                </div>
                <button type="submit" class="btn btn-primary">{{ isLogin ? 'Login' : 'Register' }}</button>
            </form>
            <div class="mt-3">
                <button class="btn btn-secondary" @click="toggleForm">{{ isLogin ? 'Switch to Register' : 'Switch to Login' }}</button>
            </div>
        </div>
    </div>
</template>
  
  <script>
  export default {
    data() {
      return {
        isLogin: true, // Start with login form
        phone: '',
        password: ''
      };
    },
    methods: {
       async handleSubmit() {
        if (this.isLogin) {
          try {
          const response = await fetch('http://localhost:3000/api/auth/login', {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json'
            },
            body: JSON.stringify({
              phone: this.phone,
              password: this.password
            })
          });
          const data = await response.json();
          if (response.ok) {
            localStorage.setItem('token', JSON.stringify(data.token));
            alert(data.message)
            this.$router.push('/profile');
            console.log('Login successful:', data.message);
          } else {
            alert(data.error)
            console.error('Login failed:', data.error);
          }
        } catch (error) {
          console.error('An error occurred during login:', error);
        }
         
        } else {
          // Handle registration
        try {
          const response = await fetch('http://localhost:3000/api/auth/register', {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json'
            },
            body: JSON.stringify({
              phone: this.phone,
              password: this.password
            })
          });
          const data = await response.json();
          if (response.ok) {
            alert(data.message)
            console.log('Registration successful:', data.message);
          } else {
            alert(data.error)
            console.error('Registration failed:', data.error);
          }
        } catch (error) {
          console.error('An error occurred during registration:', error);
        }
          console.log('Registering:', this.phone, this.password);
        }
      },
      toggleForm() {
        this.isLogin = !this.isLogin;
      }
    }
  };
  </script>
  
  <style>
.form-control {
  max-width: 300px; /* Set a maximum width for the input fields */
  margin: auto; /* Center the input fields horizontally */
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1); /* Add shadow */
}

.container {
  max-width: 500px; /* Set a maximum width for the container */
}

.btn {
  margin-top: 10px;
}
</style>
  
  