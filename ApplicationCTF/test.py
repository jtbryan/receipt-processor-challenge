import nacl.secret
import base64

key = base64.b64decode("dGhpc2lzYWtleXdpdGgzMmxldHRlcnNpbml0ZmV0Y2g=")
box = nacl.secret.SecretBox(key)

ciphertext = base64.b64decode("usbsgmFzQNjwMEEZVqJ6Hdy8MOJwMOiq4OxKbmluN/Ec7gS9EVVeh82JsqgcCUfxmAXXAx3xqmDz+dA8+JTtRknQwgjekROmuAbLiiYCQgddczHcu9xRGf1g0/KKeqnrcfbdIj1PUlK0W5iL22v98ZK+5gTFpPoHdqOHCJ2wesES6a18/t1C0CI7F40AnTrL2kengA0FCh4s4MY4")

plaintext = box.decrypt(ciphertext=ciphertext)
print(plaintext)

ciphertext = base64.b64decode("/KSIvzD+EfD2ohTDS8YyXWfpzHjVrp1iMQBE3XbIMi0G/BowKduAtkW9uvBwoRDicvmeFfDvaez3xw==")

plaintext = box.decrypt(ciphertext=ciphertext)
print(plaintext)