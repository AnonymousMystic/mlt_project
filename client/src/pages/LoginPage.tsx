import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
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
  const navigate = useNavigate()

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
        }, {
          withCredentials: true
        }
      )

        if (!response.status) {
          alert(`Fetch Error: ${response.data.message}`)
        } else {
          alert(`Message: ${response.data.message}`)
          setMessage(response.data.message)
          setError('')
          navigate('/calendar')
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
