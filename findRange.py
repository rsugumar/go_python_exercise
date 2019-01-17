import sys, getopt
from ctypes import *

class GoString(Structure):
    _fields_ = [("p", c_char_p), ("n", c_longlong)]

def main(argv):
    config_file = ''
    input_amt = ''
    try:
        opts, _ = getopt.getopt(argv, "hc:i:", ["configfile="])
    except getopt.GetoptError:
        print('findRange.py -c <config_json_file>')
        sys.exit(2)
    
    for opt, arg in opts:
        if opt == '-h':
            print('findRange.py -c <config_json_file> -i <input_amount_within_quotes>')
            sys.exit()
        elif opt in ("-c", "--configfile"):
            config_file = arg
        elif opt in ("-i", "--input"):
            input_amt = arg
    
    if not config_file or not input_amt:
        print('findRange.py -c <config_json_file> -i <input_amount_within_quotes>')
        sys.exit(2)

    print('Config: ', config_file)
    print('Input: ', input_amt)

    lib = CDLL('./libparseAmount.so')
    lib.processInput.argTypes = [GoString, GoString]
    lib.processInput.restype = None
    config_file_go_str = GoString(str.encode(config_file), len(config_file))
    amt_go_str = GoString(str.encode(input_amt), len(input_amt))
    lib.processInput(config_file_go_str, amt_go_str)

if __name__ == '__main__':
    main(sys.argv[1:])