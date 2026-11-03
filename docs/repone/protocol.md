# RepOne Tether Protocol

The teather runs a Bluetooth Low Energy (BLE) protocol to publish its data.

Here's the list of published services and charactertistics, most of which are
custom.

```
Found service: a5183278-ca65-45b7-b6c3-a68552f2026d
	Found characteristic: a5183278-ca65-45b7-b6c3-a68552f20273
	Found characteristic: a5183278-ca65-45b7-b6c3-a68552f20274
	Found characteristic: a5183278-ca65-45b7-b6c3-a68552f20275
	Found characteristic: a5183278-ca65-45b7-b6c3-a68552f20276
	Found characteristic: a5183278-ca65-45b7-b6c3-a68552f20277
	Found characteristic: a5183278-ca65-45b7-b6c3-a68552f20278
	Found characteristic: a5183278-ca65-45b7-b6c3-a68552f20279
	Found characteristic: a5183278-ca65-45b7-b6c3-a68552f2027a
	Found characteristic: a5183278-ca65-45b7-b6c3-a68552f2027b
	Found characteristic: a5183278-ca65-45b7-b6c3-a68552f2027c
	Found characteristic: a5183278-ca65-45b7-b6c3-a68552f2027d
	Found characteristic: a5183278-ca65-45b7-b6c3-a68552f2027e
	Found characteristic: a5183278-ca65-45b7-b6c3-a68552f2027f
	Found characteristic: a5183278-ca65-45b7-b6c3-a68552f20280
	Found characteristic: a5183278-ca65-45b7-b6c3-a68552f20281
	Found characteristic: a5183278-ca65-45b7-b6c3-a68552f20282
	Found characteristic: a5183278-ca65-45b7-b6c3-a68552f20283
Found service: a5183278-ca65-45b7-b6c3-a68552f4026d
	Found characteristic: a5183278-ca65-45b7-b6c3-a68552f4026e
	Found characteristic: a5183278-ca65-45b7-b6c3-a68552f4026f
Found service: a5183278-ca65-45b7-b6c3-a68552f3026d
	Found characteristic: a5183278-ca65-45b7-b6c3-a68552f3026e
	Found characteristic: a5183278-ca65-45b7-b6c3-a68552f3026f
	Found characteristic: a5183278-ca65-45b7-b6c3-a68552f30270

# Custom UUID of Nordic Semiconductor ASA per 
# https://gist.github.com/ariccio/2882a435c79da28ba6035a14c5c65f22
Found service: 0000fe59-0000-1000-8000-00805f9b34fb
	Found characteristic: 8ec90001-f315-4f60-9fb8-838830daea50
	Found characteristic: 8ec90002-f315-4f60-9fb8-838830daea50

## Batery Service
Found service: 0000180f-0000-1000-8000-00805f9b34fb
	Found characteristic: 00002a19-0000-1000-8000-00805f9b34fb
```

# Service a5183278-ca65-45b7-b6c3-a68552f2026d

this _seems_ to be the velocity service itself.

## Characteristic a5183278-ca65-45b7-b6c3-a68552f20274

Seems to be a rep counter, but it's a total one, so the amount of reps coming
as a total.

Values look like:

```
0200000000000000000000000000
0300000000000000000000000000
0400000000000000000000000000
```

Values are published from this characteristic at the start of each rep. This
value is set at the start and doesn't "publish" a value on the first rep.

## Characteristic a5183278-ca65-45b7-b6c3-a68552f20273

Seems to be the velocity measurement charactertistic? These values look like:

```
0b0081003002d801250341014c03c300ffffffff2902
0a008000d702f20152045b01ad02cd00ffffffffe101
08007e0021034802e1045f01da02f100ffffffffbf01
07007d00080318025a04cd00b302dd00ffffffff3e01
06007c00b6022502da03cd001803e200ffffffff5201
05007b00dd022902eb03cf00f302e400ffffffff3301
04007a002d02ce01be03cd003f03bf00ffffffff8e01
```

clearly some patterns here so not just integer values, the first part byte
seems to be some sort of counter-like thing so possibly a correleation with the
rep counter characteristic.
