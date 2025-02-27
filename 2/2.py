from flask import Flask, request, jsonify
import requests

app = Flask(__name__)

@app.route('/pokemon', methods=['POST'])
def get_pokemon_data():
    try:
        # รับค่า id จาก request
        data = request.get_json()
        pokemon_id = str(data.get('id')).strip()

        if not pokemon_id.isdigit():
            return jsonify({"error": "Invalid id parameter, must be a number"}), 400

        # URLs ของ PokeAPI
        pokemon_url = f"https://pokeapi.co/api/v2/pokemon/{pokemon_id}/"
        pokemon_form_url = f"https://pokeapi.co/api/v2/pokemon-form/{pokemon_id}/"

        headers = {"User-Agent": "Mozilla/5.0"}

        # ส่ง request ไปยัง PokeAPI
        pokemon_response = requests.get(pokemon_url, headers=headers)
        pokemon_form_response = requests.get(pokemon_form_url, headers=headers)

        # ตรวจสอบว่า API ตอบกลับสำเร็จหรือไม่
        if pokemon_response.status_code != 200 or pokemon_form_response.status_code != 200:
            return jsonify({"error": "Failed to fetch data from PokeAPI"}), 500

        # แปลงข้อมูลเป็น JSON
        pokemon_data = pokemon_response.json()
        pokemon_form_data = pokemon_form_response.json()

        # ดึงเฉพาะข้อมูลที่ต้องการ
        result = {
            "stats": pokemon_data.get("stats", []),
            "name": pokemon_form_data.get("name", ""),
            "prites": pokemon_form_data.get("sprites", {})
        }

        return jsonify(result)

    except requests.exceptions.RequestException as e:
        return jsonify({"error": f"Request failed: {str(e)}"}), 500
    except Exception as e:
        return jsonify({"error": str(e)}), 500

if __name__ == '__main__':
    app.run(debug=True)
