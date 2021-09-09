import serial
import time
import pygame

pygame.init()

screen = pygame.display.set_mode([500,500])

ser = serial.Serial("/dev/ttyUSB1")
print("waiting")
time.sleep(4)
print("go")

def writeToRobot(payload):
    print(payload)
    ser.write(payload)

def execute(batch):
    numMotor = len(batch.movementsByMotor)

    packet = bytearray()
    packet.append(numMotor)
    for motorId in batch.movementsByMotor.keys():
        packet.append(motorId)
        packet.append(len(batch.movementsByMotor[motorId]))
        for commandId, movement in enumerate(batch.movementsByMotor[motorId]):
            print(commandId)
            packet.append(commandId)
            print(movement[0])
            movementBytes = movement[0].to_bytes(2, 'big', signed=True) 
            for b in movementBytes:
                packet.append(b)
            durationBytes = movement[1].to_bytes(2, 'big', signed=True) 
            for b in durationBytes:
                packet.append(b)
    writeToRobot(packet)


    
class Batch:
    def __init__(self):
        self.movementsByMotor = {}
        

    def add_movement(self, motor, movement, duration):
        if motor in self.movementsByMotor:
            self.movementsByMotor[motor].append((movement, duration))
        else:
            self.movementsByMotor[motor] = [(movement, duration)]


claw = 0
secondary = 1
z = 2
primary = 3

history = []
inverse_history = []
while(True):
    for event in pygame.event.get():
        if event.type == pygame.KEYDOWN:
            if event.key == pygame.K_a:
                batch = Batch()
                batch.add_movement(z, 5, 30)
                history.append((batch, 30))
                execute(batch)
                inverse = Batch()
                inverse.add_movement(z, -5, 30)
                inverse_history.append((inverse,30))
                time.sleep(0.03)
            if event.key == pygame.K_d:
                batch = Batch()
                batch.add_movement(z, -5, 30)
                history.append((batch,30))
                execute(batch)
                inverse = Batch()
                inverse.add_movement(z, 5, 30)
                inverse_history.append((inverse, 30))
                time.sleep(0.03)
            if event.key == pygame.K_w:
                batch = Batch()
                batch.add_movement(primary, 10, 10)
                history.append((batch, 10))
                execute(batch)
                inverse = Batch()
                inverse.add_movement(primary, -10, 10)
                inverse_history.append((inverse, 10))
                time.sleep(0.01)
            if event.key == pygame.K_s:
                batch = Batch()
                batch.add_movement(primary, -10, 10)
                history.append((batch, 10))
                execute(batch)
                inverse = Batch()
                inverse.add_movement(primary, 10, 10)
                inverse_history.append((inverse, 10))
                time.sleep(0.01)
            if event.key == pygame.K_e:
                batch = Batch()
                batch.add_movement(secondary, 5, 30)
                history.append((batch,30))
                execute(batch)
                inverse = Batch()
                inverse.add_movement(z, 5, 30)
                inverse_history.append((inverse, 30))
                time.sleep(0.03)
            if event.key == pygame.K_q:
                batch = Batch()
                batch.add_movement(secondary, -5, 30)
                history.append((batch,30))
                inverse = Batch()
                inverse.add_movement(secondary, 5, 30)
                inverse_history.append((inverse,30))
                execute(batch)
                time.sleep(0.03)
            if event.key == pygame.K_o:
                batch = Batch()
                batch.add_movement(claw, -5, 30)
                history.append((batch, 30))
                inverse = Batch()
                inverse.add_movement(claw, 5, 30)
                inverse_history.append((inverse,30))
                execute(batch)
                time.sleep(0.03)
            if event.key == pygame.K_c:
                batch = Batch()
                batch.add_movement(claw, 5, 30)
                history.append((batch, 30))
                inverse = Batch()
                inverse.add_movement(claw, -5, 30)
                inverse_history.append((inverse, 30))
                execute(batch)
                time.sleep(0.03)

            if event.key == pygame.K_r:
                for cmd in inverse_history[::-1]:
                    execute(cmd[0])
                    time.sleep(cmd[1]/100)
                for cmd in history:
                    execute(cmd[0])
                    time.sleep(cmd[1]/100)




while True:
    pass
print("killing bot")
