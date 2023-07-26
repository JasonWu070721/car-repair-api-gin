import requests


# Customers
host = "http://127.0.0.1:3000/api/v1"

id = 0
payload = {'name': 'name', 'license_plate': 'license_plate',
           'car_color': 'car_color', 'car_year': 'car_year'}

r = requests.post(f'{host}/customers', json=payload)
if (r.status_code == 201):
    id = r.json()['ID']


r = requests.get(f'{host}/customers')
if (r.status_code == 200):
    print(r.json())

if id != 0:
    r = requests.get(f'{host}/customers/{id}')
    if (r.status_code == 200):
        print(r.json())

    r = requests.delete(f'{host}/customers/{id}')
    if (r.status_code == 204):
        print(r.json())
