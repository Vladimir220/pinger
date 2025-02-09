const { useState, useEffect } = React;

function App() {
    const [ipList, setIpList] = useState([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        fetchIpList(); 

        const intervalId = setInterval(fetchIpList, 5000);

        return () => clearInterval(intervalId); 
    }, []);

    const fetchIpList = async () => {
        try {
            const response = await axios.get('http://localhost:1234/read_all_statuses');
            

            const formattedData = response.data.map(host => ({
                ...host,
                last_success_date: formatDate(host.last_success_date)
            }));
            setIpList(formattedData);
        } catch (error) {
            console.error('Error fetching IP list:', error);
        } finally {
            setLoading(false);
        }
    };

    const formatDate = (dateString) => {
        if (!dateString) return '';
    
        const date = new Date(dateString);
        const utcOffset = date.getTimezoneOffset() * 60000;
        const adjustedDate = new Date(date.getTime() + utcOffset);
    
        const year = adjustedDate.getFullYear();
        const month = String(adjustedDate.getMonth() + 1).padStart(2, '0');
        const day = String(adjustedDate.getDate()).padStart(2, '0');
        const hours = String(adjustedDate.getHours()).padStart(2, '0');
        const minutes = String(adjustedDate.getMinutes()).padStart(2, '0');
        const seconds = String(adjustedDate.getSeconds()).padStart(2, '0');
    
        return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
    };

    return (
        <div className="container">
            <h1 className="mt-4">Список отслеживаемых хостов</h1>
            {loading ? (
                <p>Загрузка...</p>
            ) : (
                <table className="table table-striped mt-4">
                    <thead>
                        <tr>
                            <th>IP Адрес</th>
                            <th>Время Пинга (мс)</th>
                            <th>Дата Последней Успешной Попытки</th>
                        </tr>
                    </thead>
                    <tbody>
                        {ipList.map((host, index) => (
                            <tr key={index}>
                                <td>{host.ip}</td>
                                <td>{host.ping_time_ms}</td>
                                <td>{host.last_success_date}</td> {/* Now it's formatted */}
                            </tr>
                        ))}
                    </tbody>
                </table>
            )}
        </div>
    );
}

window.App = App;