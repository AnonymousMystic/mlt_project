import React, { useState } from 'react'
import { useNavigate } from 'react-router-dom';
import axios from 'axios'

const RegistrationForm = () => {
  const [error, setError] = useState<string>('');
  const [email, setEmail] = useState<string>('');
  const [password, setPassword] = useState<string>('');

  const navigate = useNavigate()

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    (async () => {
        try {
            const response = await axios.post('http://localhost:8080/api/auth/register', { 
                "username": email, 
                "password": password 
            }, {
                withCredentials: true
            })

            // check for valid registration
            if (!response.status) {
                alert("error registering")
                setError("Error registering")
            } else {
                setError('')
                navigate('/calendar')
            }
            setEmail('');
            setPassword('');
        } catch (err) {
            console.log("Encountered Registration Error: ", err)
            setError("Error registering")
        }

    })()
  };

  return (
    <div>
      <h2>Register</h2>
      <form onSubmit={handleSubmit}>
        <input
          type="email"
          placeholder="Email"
          required
          value={email}
          onChange={(e: React.ChangeEvent<HTMLInputElement>) => setEmail(e.target.value)}
        />

        <input
          type="password"
          placeholder="Password"
          required
          value={password}
          onChange={(e: React.ChangeEvent<HTMLInputElement>) => setPassword(e.target.value)}
        />

        <button type="submit">
          Register
        </button>
      </form>
      <div>
        {error}
      </div>
    </div>
  );
};

export default RegistrationForm