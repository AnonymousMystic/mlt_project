import 'react'
import LoginForm from './pages/LoginPage'
import CalendarPage from './pages/CalendarPage'
import { Routes, Route } from 'react-router-dom'
import LandingPage from './pages/LandingPage'
import RegistrationForm from './pages/ResgiterPage'

function App() {
  return (
    <Routes>
      <Route path="/" element={<LandingPage/>} />
      <Route path="/calendar" element={<CalendarPage />} />
      <Route path="/login" element={<LoginForm />} />
      <Route path="/register" element={<RegistrationForm />} />
    </Routes>
  )
}

export default App
