import React, { useState, useEffect } from 'react';
import axios from 'axios';

interface LoginFormState {
  email: string;
  password: string;
}

const LoginForm: React.FC = () => {
  const [formData, setFormData] = useState<LoginFormState>({
    email: '',
    password: '',
  });

  const [error, setError] = useState<string>('');
  const [message, setMessage] = useState<string>('')

  // attempt to access profile without logging in
  useEffect(() => {
    async function retrieveProfile() {
      try {
        const response = await axios.get("http://localhost:8080/api/auth/profile")

        if (response.status) {
          alert(response.data.message)
          console.log("Already authenticated, redirecting to login page")
        }

      } catch {
        console.log("Not authenticated, redirect to login page")
      }
    }

    retrieveProfile()
  }, [])

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();

    const { email, password } = formData;

    if (!email || !password) {
      setError('Please fill in all fields');
      return;
    }

    // Simulated login
    if (email === "user@example.com" && password === "password123") {  
      (async () => {
        const response = await axios.post('http://localhost:8080/api/auth/login', { 
          "username": email, 
          "password": password 
        })

        if (!response.status) {
          alert(`Fetch Error: ${response.data.message}`)
        } else {
          alert(`Message: ${response.data.message}`)
          setMessage(response.data.message)
          setError('')
        }
      })()
    } else {
      setError('Invalid email or password');
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <h2>Login</h2>
      <input
        type="email"
        name="email"
        placeholder="Email"
        value={formData.email}
        onChange={handleChange}
      />
      <input
        type="password"
        name="password"
        placeholder="Password"
        value={formData.password}
        onChange={handleChange}
      />
      {error && <p>{error}</p>}
      <button type="submit">Login</button>
      {message && <p>{message}</p>}
    </form>
  );
};

export default LoginForm;
