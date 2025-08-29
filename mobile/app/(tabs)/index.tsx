import { Pressable, SafeAreaView, StyleSheet, Text } from "react-native";

export default function HomeScreen() {
	const BASE_URL = "https://noizyboy.onrender.com/";
	return (
		<SafeAreaView style={[styles.container]}>
			<Text style={[styles.titleText]}>Noizy Boy</Text>
			<Pressable
				android_ripple={{ color: "lightblue", borderless: false }}
				onPress={() => {
					fetch(BASE_URL + "api/v1/make-some-noise", {
						method: "POST",
						body: JSON.stringify({
							push_token: "expo_push_token",
							name: "user_name",
						}),
					});
				}}
				style={[styles.btn]}
			>
				<Text style={[styles.btnText]}>Talk</Text>
			</Pressable>
		</SafeAreaView>
	);
}

const styles = StyleSheet.create({
	container: {
		flex: 1,
		gap: 8,
		alignItems: "center",
		justifyContent: "center",
		paddingHorizontal: 16,
	},
	titleText: {
		color: "cyan",
		fontWeight: "600",
		fontSize: 24,
	},
	btn: {
		paddingInline: 16,
		paddingBlock: 24,
		backgroundColor: "blue",
		minWidth: "50%",
		borderRadius: 8,
		alignItems: "center",
		justifyContent: "center",
		marginTop: 12,
	},
	btnText: {
		color: "white",
		fontSize: 16,
		fontWeight: "400",
	},
});
