import { useState } from "react";
import { View, TextInput, Text } from "react-native";

const API_BASE = "http://127.0.0.1:9090";

export default function SearchScreen() {
  const [query, setQuery] = useState("");
  const [data, setData] = useState([]);

  const search = (text) => {
    setQuery(text);

    if (!text) {
      setData([]);
      return;
    }

    fetch(`${API_BASE}/search?q=${text}`)
      .then((res) => res.json())
      .then((json) => {
        console.log("SEARCH DATA:", json);
        setData(json);
      })
      .catch((err) => console.error("ERROR:", err));
  };

  return (
    <View style={{ padding: 20 }}>
      <TextInput
        placeholder="Search username"
        value={query}
        onChangeText={search}
        style={{ padding: 10, borderWidth: 1, marginBottom: 10 }}
      />

      {data.map((item, index) => (
        <View
          key={item.username + index}
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
