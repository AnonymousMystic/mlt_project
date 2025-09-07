import 'react'
import LoginForm from './components/LoginPage'
import CalendarPage from './components/CalendarPage'
import { Routes, Route } from 'react-router-dom'
import LandingPage from './components/LandingPage'

function App() {
  return (
    <Routes>
      <Route path="/" element={<LandingPage/>} />
      <Route path="/calendar" element={<CalendarPage />} />
      <Route path="/login" element={<LoginForm />} />
    </Routes>
  )
}

export default App
