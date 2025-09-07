import { useState } from 'react'

// calendar page component for assigning users
const CalendarPage: React.FC = () => {
    const daysOfWeek = ['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat'];
    const today = new Date();
    const [currentDate, setCurrentDate] = useState(new Date(today.getFullYear(), today.getMonth(), 1));

    const year = currentDate.getFullYear();
    const month = currentDate.getMonth(); // 0-indexed

    const firstDayOfMonth = new Date(year, month, 1);
    const lastDayOfMonth = new Date(year, month + 1, 0);

    const daysInMonth = lastDayOfMonth.getDate();
    const startDay = firstDayOfMonth.getDay(); // 0 = Sunday, 6 = Saturday

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

    // function calendarRetrieval() {
    //     // Retrieves calendar
    // }

    // function createAppointment() {
    //     // Creates calendar appointment
    // }

    return (
        <div>
        <div>
            <button onClick={handlePrevMonth}>Previous</button>
            <span>{currentDate.toLocaleString('default', { month: 'long' })} {year}</span>
            <button onClick={handleNextMonth}>Next</button>
        </div>

        <div style={{ display: 'grid', gridTemplateColumns: 'repeat(7, 1fr)' }}>
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

