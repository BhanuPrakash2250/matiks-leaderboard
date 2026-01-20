import { useEffect, useState } from "react";
import { View, Text } from "react-native";

const API_BASE = "http://127.0.0.1:9090";

export default function LeaderboardScreen() {
  const [data, setData] = useState([]);

  useEffect(() => {
    fetch(`${API_BASE}/leaderboard?limit=10`)
      .then((res) => res.json())
      .then((json) => {
        console.log("LEADERBOARD DATA:", json);
        setData(json);
      })
      .catch((err) => console.error("ERROR:", err));
  }, []);

  return (
    <View style={{ padding: 20 }}>
      <Text style={{ fontWeight: "bold", marginBottom: 10 }}>
        Rank | Username | Rating
      </Text>

      {data.map((item, index) => (
        <View
          key={item.id || index}
          style={{ flexDirection: "row", paddingVertical: 6 }}
        >
          <Text style={{ width: 60 }}>{item.rank}</Text>
          <Text style={{ width: 160 }}>{item.username}</Text>
          <Text>{item.rating}</Text>
        </View>
      ))}
    </View>
  );
}
