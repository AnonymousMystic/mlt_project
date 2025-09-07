import { useState, useEffect } from 'react'
import { useNavigate } from 'react-router-dom'
import axios from 'axios'

const LandingPage: React.FC = () => { 
    const [loading, setLoading] = useState('Loading...')
    const navigate = useNavigate()

    // attempt to access profile without logging in
    useEffect(() => {
        async function retrieveCalendar() {
        try {
            const response = await axios.get("http://localhost:8080/api/user/profile", 
                {
                    withCredentials: true
                }
            )

            if (response.status) {
                console.log("Already authenticated, redirecting to profile page")
                setLoading("Redirecting to login")
                navigate('/calendar')
            }

        } catch {
            console.log("Not authenticated, redirect to login page")
            setLoading("Redirecting to calendar")
            navigate('/login')
        }
    }

    retrieveCalendar()
  }, [])
  
    return (
        <>
            <div className="text-blue-500">
                {loading}
            </div>
        </>
    )
}

export default LandingPage
