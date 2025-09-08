import { useState, useEffect } from 'react'
import axios from 'axios'
import { useNavigate } from 'react-router-dom'

// calendar page component for assigning users
const CalendarPage: React.FC = () => {
    const daysOfWeek = ['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat'];
    const today = new Date();
    const [currentDate, setCurrentDate] = useState(new Date(today.getFullYear(), today.getMonth(), 1));
    const navigate = useNavigate()

    const year = currentDate.getFullYear();
    const month = currentDate.getMonth(); // 0-indexed

    const firstDayOfMonth = new Date(year, month, 1);
    const lastDayOfMonth = new Date(year, month + 1, 0);

    const daysInMonth = lastDayOfMonth.getDate();
    const startDay = firstDayOfMonth.getDay(); // 0 = Sunday, 6 = Saturday

    // attempt to access profile with tokens
    useEffect(() => {
        async function retrieveCalendar() {
            const response = await axios.get("http://localhost:8080/api/user/profile", 
                {
                    withCredentials: true
                }
            )

            if (!response.status) {
                console.log("Not authenticated, redirect to login page")
                navigate('/login')
            }
        }

        retrieveCalendar()
    }, [])

    const generateCalendarDays = (): (number | null)[] => {
        const days: (number | null)[] = [];

        // Padding for the first row (nulls before the 1st of the month)
        for (let i = 0; i < startDay; i++) {
            days.push(null);
        }

        // Add actual days
        for (let day = 1; day <= daysInMonth; day++) {
            days.push(day);
        }

        return days;
    };

    const handlePrevMonth = () => {
        setCurrentDate(prev => new Date(prev.getFullYear(), prev.getMonth() - 1, 1));
    };

    const handleNextMonth = () => {
        setCurrentDate(prev => new Date(prev.getFullYear(), prev.getMonth() + 1, 1));
    };

    const calendarDays = generateCalendarDays();

    const handleLogout = () => {
        (async () => {
            try {
                const response = await axios.post('http://localhost:8080/api/auth/logout', null, {
                        withCredentials: true
                    }
                )

                if (response.status) {
                    console.log("Log out successful")
                    navigate('/login')
                }
            } catch (err) {
                console.log("Logout error:", err)
            }
        })()
    }

    // TODO: API calls for event and appointment scheduling
    // function calendarRetrieval() {
    //     // Retrieves calendar
    // }

    // function createAppointment() {
    //     // Creates calendar appointment
    // }

    return (
        <div>
            <div className='flex'>
                <button onClick={handlePrevMonth}>Previous</button>
                <span>{currentDate.toLocaleString('default', { month: 'long' })} {year}</span>
                <button onClick={handleNextMonth}>Next</button>
                <button className="ml-auto" onClick={handleLogout}>Logout</button>
            </div>

            <div className="grid grid-cols-7">
                {daysOfWeek.map(day => (
                <div key={day}><strong>{day}</strong></div>
                ))}

                {calendarDays.map((day, index) => (
                <div key={index}>{day !== null ? day : ''}</div>
                ))}
            </div>
        </div>
    );
};

export default CalendarPage

