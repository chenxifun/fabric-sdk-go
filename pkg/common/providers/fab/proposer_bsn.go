/**
 * @Author: Gao Chenxi
 * @Description:
 * @File:  proposer_bsn
 * @Version: 1.0.0
 * @Date: 2020/4/8 15:18
 */

package fab

import (
	pb "github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/protos/peer"
)

type RequestProposal struct {
	TransactionProposal *TransactionProposal
	SignProposal        *pb.SignedProposal
}
